package main

import (
	"example/hello/app"
	"example/hello/config"
	"example/hello/server"

	"github.com/urfave/cli/v2"

	"log"
	"os"
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
	dependency := app.Init(conf)
	server.StartAPIServer(dependency)
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
