package cmd

import (
	"github.com/spf13/cobra"
	"os/exec"
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
			// run server at cmd/server/serverapp.go
			server := exec.Command("go", "run", "cmd/server/serverapp.go")
			// pass stdin, stdout, and stderr to child process
			server.Stdin = cmd.InOrStdin()
			server.Stdout = cmd.OutOrStdout()
			server.Stderr = cmd.ErrOrStderr()
			err := server.Run()
			if err != nil {
				cmd.PrintErrf("error running server: %s", err)
			}
		},
	}
)
