package cmd

import (
	"net/http"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "GET",
	Aliases: []string{"get"},
	Short:   "GET request",
	Args:    cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		// parse the arg as a url
		urlStr := args[0]
		url, err := parseURL(urlStr)
		if err != nil {
			return err
		}

		resp, err := http.Get(url.String())
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		err = printResp(c, resp, "html")
		if err != nil {
			return err
		}

		return nil
	},
}
