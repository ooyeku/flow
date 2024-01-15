package cmd

import (
	"github.com/spf13/cobra"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(serverCommand)
}

func executeServer(cmd *cobra.Command) {
	// run server at cmd/server/serverapp.go
	server := exec.Command("go", "run", "cmd/server/serverapp.go")
	// pass stdin, stdout, and stderr to child process
	server.Stdin = cmd.InOrStdin()
	server.Stdout = cmd.OutOrStdout()
	server.Stderr = cmd.ErrOrStderr()

	handleError(cmd, runServer(server))
}

func runServer(server *exec.Cmd) error {
	return server.Run()
}

func handleError(cmd *cobra.Command, err error) {
	if err != nil {
		cmd.PrintErrf("error running server: %s", err)
	}
}

var (
	serverCommand = &cobra.Command{
		Use:   "server",
		Short: "run workflow in server mode",
		Long:  "run workflow in server mode",
		Run: func(cmd *cobra.Command, args []string) {
			executeServer(cmd)
		},
	}
)
