package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "http",
		Usage: "a http client implemented in go",
		Commands: []*cli.Command{
			{
				Name:      "GET",
				Aliases:   []string{"get"},
				Usage:     "GET request",
				UsageText: "send a request using GET method",
				Action:    getAction,
			},
			{
				Name:      "POST",
				Aliases:   []string{"post"},
				Usage:     "POST request",
				UsageText: "send a request using POST method",
				Action:    postAction,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			switch cmd.Args().Len() {
			case 0:
				return ErrNoArg
			case 1:
				return getAction(ctx, cmd)
			default:
				return postAction(ctx, cmd)
			}
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
