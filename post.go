package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func post(c *cli.Context) error {
	// throw error if no args provided
	if c.Args().Len() == 0 {
		return errNoArgs
	}

	urlStr := c.Args().First()
	u, e := processUrl(urlStr)
	if e != nil {
		return e
	}

	data := make(map[string]interface{})
	for i := 1; i < c.Args().Len(); i++ {
		arg := c.Args().Get(i)
		idx := strings.IndexRune(arg, '=')
		k := arg[:idx]
		v := arg[idx+1:]
		if v == "true" {
			data[k] = true
		} else if v == "false" {
			data[k] = false
		} else {
			num, e := strconv.Atoi(v)
			if e == nil {
				data[k] = num
			} else {
				data[k] = v
			}
		}
	}

	jsonStr, e := json.Marshal(data)
	if e != nil {
		return e
	}

	r, e := http.Post(u.String(), "application/json", bytes.NewReader(jsonStr))
	if e != nil {
		return e
	}

	e = printResp(r, "json")
	if e != nil {
		return e
	}
	r.Body.Close()

	return nil
}
