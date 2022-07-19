package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var errNoArgs = errors.New("no args provided")

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
			{
				Name:      "POST",
				Aliases:   []string{"post"},
				Usage:     "POST request",
				UsageText: "send a request using POST method",
				Action:    post,
			},
		},
		Action: func(c *cli.Context) error {
			switch c.Args().Len() {
			case 0:
				return errNoArgs
			case 1:
				return get(c)
			default:
				return post(c)
			}
		},
	}

	e := app.Run(os.Args)
	if e != nil {
		log.Fatal(e)
	}
}
