package main

import (
	"context"
	"net/http"

	"github.com/urfave/cli/v3"
)

func getAction(ctx context.Context, cmd *cli.Command) error {
	// throw error if no args provided
	if cmd.Args().Len() == 0 {
		return ErrNoArg
	}

	// parse the arg as a url
	urlStr := cmd.Args().First()
	url, err := parseUrl(urlStr)
	if err != nil {
		return err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = printResp(resp, "html")
	if err != nil {
		return err
	}

	return nil
}
