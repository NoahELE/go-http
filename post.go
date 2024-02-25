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
	u, e := parseUrl(urlStr)
	if e != nil {
		return e
	}

	data := make(map[string]interface{})
	for i := 1; i < cmd.Args().Len(); i++ {
		arg := cmd.Args().Get(i)
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
	defer r.Body.Close()

	e = printResp(r, "json")
	if e != nil {
		return e
	}

	return nil
}
