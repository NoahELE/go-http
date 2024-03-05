package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/urfave/cli/v3"
)

func postAction(ctx context.Context, cmd *cli.Command) error {
	// throw error if no args provided
	if cmd.Args().Len() == 0 {
		return ErrNoArg
	}

	urlStr := cmd.Args().First()
	url, err := parseUrl(urlStr)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	for i := 1; i < cmd.Args().Len(); i++ {
		arg := cmd.Args().Get(i)
		if strings.Contains(arg, ":=") {
			idx := strings.Index(arg, ":=")
			k := arg[:idx]
			v := arg[idx+2:]
			if v == "true" {
				data[k] = true
			} else if v == "false" {
				data[k] = false
			} else {
				n, err := strconv.Atoi(v)
				if err != nil {
					data[k] = v
				}
				data[k] = n
			}
		} else if strings.Contains(arg, "=") {
			idx := strings.Index(arg, "=")
			k := arg[:idx]
			v := arg[idx+1:]
			data[k] = v
		} else {
			return ErrInvalidArg
		}
	}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	r, err := http.Post(url.String(), "application/json", bytes.NewReader(jsonStr))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = printResp(r, "json")
	if err != nil {
		return err
	}

	return nil
}
