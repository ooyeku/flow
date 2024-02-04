package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(CliCmd)
}

var CliCmd = &cobra.Command{
	Use:   "cli",
	Short: "run flow in cli mode",
	Long:  `run flow in cli mode`,
	Run: func(cmd *cobra.Command, args []string) {
		// run cli loop
		cli := exec.Command("go", "run", "cmd/cli/cliapp.go")
		cli.Stdin = os.Stdin
		cli.Stdout = os.Stdout
		cli.Stderr = os.Stderr
		err := cli.Run()
		if err != nil {
			cmd.PrintErrf("error running cli: %s", err)
		}
	},
}
