package server

import (
	"errors"

	"github.com/sam701/file-publisher/cleanup"
	"github.com/sam701/file-publisher/config"
	"github.com/urfave/cli"
)

var (
	addrFlag = cli.StringFlag{
		Name:   "addr",
		Usage:  "Address to start server on",
		Value:  "localhost:9000",
		EnvVar: "ADDR",
	}
	dataDirFlag = cli.StringFlag{
		Name:  "data-dir",
		Usage: "Path to the data directory",
	}
	baseUrlFlag = cli.StringFlag{
		Name:  "base-url",
		Usage: "Server's base URL used for sharing",
	}

	Command = cli.Command{
		Name:        "server",
		Description: "Start the file server",
		Action:      serverAction,
		Flags: []cli.Flag{
			addrFlag,
			dataDirFlag,
			baseUrlFlag,
		},
	}
)

func serverAction(ctx *cli.Context) error {
	dd := ctx.String(dataDirFlag.Name)
	if dd == "" {
		return errors.New("No data dir was specified")
	}
	bu := ctx.String(baseUrlFlag.Name)

	config.Read(dd, bu)

	go cleanup.Run()
	startServer(ctx.String(addrFlag.Name))
	return nil
}
