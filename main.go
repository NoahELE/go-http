package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "http",
		Usage: "a http client implemented in go",
		Commands: []*cli.Command{
			{
				Name:      "GET",
				Aliases:   []string{"get"},
				Usage:     "GET request",
				UsageText: "send a request using GET method",
				Action:    get,
			},
		},
		Action: get,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
