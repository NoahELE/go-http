package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "http",
	Short: "a http client implemented in go",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errNoMethod
	},
}

func initRootCmd() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(postCmd)
}

func Execute() {
	initRootCmd()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
