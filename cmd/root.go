package cmd

import (
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
)

func Execute() error {

	return rootCmd.Execute()
}
