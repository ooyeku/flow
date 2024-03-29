package cmd

import (
	"github.com/spf13/cobra"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(serverCommand)
}

/*
Right now if the server is running, cli command will work but not connect to db.  If the cli
is running, the server command will not work.  Need to figure out how to run both at the same time, regardless
of which one is started first.
*/

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
		Short: "run flow in server mode",
		Long:  "run flow in server mode",
		Run: func(cmd *cobra.Command, args []string) {
			executeServer(cmd)
		},
	}
)
