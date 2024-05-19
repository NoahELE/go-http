package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:     "POST",
	Aliases: []string{"post"},
	Short:   "POST request",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		urlStr := args[0]
		url, err := parseURL(urlStr)
		if err != nil {
			return err
		}

		data := make(map[string]interface{})
		for i := 1; i < len(args); i++ {
			arg := args[i]
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
				return errInvalidArg
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

		err = printResp(c, r, "json")
		if err != nil {
			return err
		}

		return nil
	},
}
