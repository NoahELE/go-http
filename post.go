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

	url := c.Args().First()
	url = processUrl(url)

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
			num, err := strconv.Atoi(v)
			if err == nil {
				data[k] = num
			} else {
				data[k] = v
			}
		}
	}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonStr))
	if err != nil {
		return err
	}

	err = printResp(resp, "json")
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}
