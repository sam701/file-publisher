package main

import (
	"log"
	"os"

	"github.com/sam701/file-publisher/client"
	"github.com/sam701/file-publisher/server"
	"github.com/urfave/cli"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Usage = "A file publisher tool"
	app.Commands = []cli.Command{
		server.Command,
		client.Command,
	}
	app.Run(os.Args)
}
