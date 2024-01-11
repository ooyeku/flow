package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// 	root entry
	rootCmd = &cobra.Command{
		Use:   "flow",
		Short: "A workflow and process management tool",
		Long: `Work-flow is a CLI tool for managing your workflow and processes.
				Complete documentation is available at...`,
	}

	cli = cobra.Command{
		Use:   "cli",
		Short: "run workflow in cli mode",
		Long:  `run workflow in cli mode`,
		Run: func(cmd *cobra.Command, args []string) {
			// run cli at cli/main.go
			fmt.Println("run cli at cli/main.go")
		},
	}

	server = cobra.Command{
		Use:   "server",
		Short: "run workflow in server mode",
		Long:  `run workflow in server mode`,
		Run: func(cmd *cobra.Command, args []string) {
			// run server at server/main.go
			fmt.Println("run server at server/main.go")
		},
	}
)

func init() {
	rootCmd.AddCommand(&cli)
	rootCmd.AddCommand(&server)
}

func Execute() error {
	return rootCmd.Execute()
}
