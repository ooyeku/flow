package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "run workflow in server mode",
		Long:  "run workflow in server mode",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run server at cmd/server/main.go")
		},
	}
)
