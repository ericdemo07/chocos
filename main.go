package main

import (
	"example/hello/config"
	"example/hello/server"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	Revision = "fafafaf"
)

var conf = config.Load()

var startServerCmd = &cli.Command{
	Name:    "start",
	Aliases: []string{"s"},
	Usage:   "start server",
	Action:  startServerAction,
}

func startServerAction(c *cli.Context) error {
	server.StartAPIServer(conf)
	return nil
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			startServerCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
