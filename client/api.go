package client

import (
	"fmt"
	"log"
	"net/http"

	"io"

	"bytes"
	"encoding/json"

	"github.com/sam701/file-publisher/meta"
)

func isFileExists(id string) bool {
	resp, err := http.Get(fmt.Sprintf("%s/api/files/%s/exists", serverUrl, id))
	if err != nil {
		log.Fatalln("ERROR", err)
	}

	return resp.StatusCode == 204
}

func uploadFile(id string, r io.Reader) error {
	_, err := http.Post(fmt.Sprintf("%s/api/files/%s", serverUrl, id), "application/octet-stream", r)
	return err
}

func uploadMeta(m *meta.PublishingRequest) (string, error) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(m)

	resp, err := http.Post(fmt.Sprintf("%s/api/meta", serverUrl), "application/octet-stream", bytes.NewReader(buf.Bytes()))
	if err != nil {
		return "", err
	}

	var pr meta.PublishingResponse
	json.NewDecoder(resp.Body).Decode(&pr)

	return pr.SharingURL, err
}
