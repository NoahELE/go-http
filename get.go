package main

import (
	"net/http"

	"github.com/urfave/cli/v2"
)

func get(c *cli.Context) error {
	// throw error if no args provided
	if c.Args().Len() == 0 {
		return ErrNoArg
	}

	// parse the arg as a url
	urlStr := c.Args().First()
	u, e := parseUrl(urlStr)
	if e != nil {
		return e
	}

	r, e := http.Get(u.String())
	if e != nil {
		return e
	}
	defer r.Body.Close()

	e = printResp(r, "html")
	if e != nil {
		return e
	}

	return nil
}
