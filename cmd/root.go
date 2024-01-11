package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// flags
	rootCmd = &cobra.Command{
		Use:   "work-flow",
		Short: "A workflow and process management tool",
		Long: `Work-flow is a CLI tool for managing your workflow and processes.
				Complete documentation is available at...`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
