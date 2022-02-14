package main

import (
	"net/http"

	"github.com/urfave/cli/v2"
)

func get(c *cli.Context) error {
	// throw error if no args provided
	if c.Args().Len() == 0 {
		return errNoArgs
	}

	url := c.Args().First()
	url = processUrl(url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	err = printResp(resp, "html")
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}
