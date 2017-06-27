package client

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/sam701/file-publisher/meta"
	"github.com/urfave/cli"
)

var (
	serverUrl string
)

func publishFile(ctx *cli.Context) error {
	fileName := ctx.String("file")
	serverUrl = ctx.String("server-url")
	if serverUrl == "" {
		return errors.New("Server URL is not set")
	}
	hash := fileHash(fileName)

	if !isFileExists(hash) {
		f, err := os.Open(fileName)
		if err != nil {
			log.Fatalln("ERROR: cannot open file:", err)
		}
		defer f.Close()

		err = uploadFile(hash, f)
		if err != nil {
			log.Fatalln("ERROR", err)
		}
	}

	ex := ctx.String("expire")
	if ex == "" {
		return errors.New("No expiration time was provided")
	}

	md := &meta.PublishingRequest{
		FileName:       path.Base(fileName),
		FileHash:       hash,
		ExpirationTime: parseExpiration(ex).Format(time.RFC3339),
	}

	url, err := uploadMeta(md)
	if err != nil {
		return err
	}
	fmt.Println("Share URL:", url)
	return nil
}

func fileHash(fileName string) string {
	h := sha1.New()
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("ERROR: cannot open file:", err)
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func parseExpiration(s string) time.Time {
	ix := 0
	for ; ix < len(s); ix++ {
		if s[ix] < '0' || s[ix] > '9' {
			break
		}
	}

	cnt, _ := strconv.Atoi(s[:ix])
	ti := time.Now()
	switch s[ix:] {
	case "min":
		ti = ti.Add(time.Duration(cnt) * time.Minute)
	case "h":
		ti = ti.Add(time.Duration(cnt) * time.Hour)
	case "d":
		ti = ti.AddDate(0, 0, cnt)
	case "m":
		ti = ti.AddDate(0, cnt, 0)
	case "y":
		ti = ti.AddDate(cnt, 0, 0)
	default:
		log.Fatal("Cannot parse time interval:", s[ix:])
	}

	return ti
}
