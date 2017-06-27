package client

import (
	"github.com/urfave/cli"
)

var (
	serverUrlFlag = cli.StringFlag{
		Name:  "server-url",
		Usage: "URL of the publishing server",
	}
	fileToPublishFlag = cli.StringFlag{
		Name:  "file",
		Usage: "File path to publish",
	}
	expirationFlag = cli.StringFlag{
		Name:  "expire, e",
		Usage: "Expiration time (10min, 2h, 3d, 4m, 1h)",
	}

	Command = cli.Command{
		Name:        "client",
		Description: "Run client CLI",
		Action:      publishFile,
		Flags: []cli.Flag{
			serverUrlFlag,
			fileToPublishFlag,
			expirationFlag,
		},
	}
)
