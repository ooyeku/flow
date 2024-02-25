package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var chatVersion string

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Launch a chat with the AI",
	Long: `Launch a chat with the AI using the cv1(non-streaming), sv1(streaming) version, and pv1(perplexity).

chat history will be stored in a local chat.db file and can be accessed by running scripts/getchathistory.go

The chat will run until you type 'exit' and press enter.

Example usage:
go run main.go chat --version cv1 
go run main.go chat -v cv1`,
	Run: func(cmd *cobra.Command, args []string) {
		var chat *exec.Cmd
		if chatVersion == "cv1" {
			chat = exec.Command("go", "run", "cmd/chat/cv1.go")
		} else if chatVersion == "sv1" {
			chat = exec.Command("go", "run", "cmd/chat/sv1.go")
		} else if chatVersion == "pv1" {
			chat = exec.Command("go", "run", "cmd/chat/pv1.go")
		} else {
			cmd.PrintErrf("invalid chat version: %s", chatVersion)
			return
		}
		chat.Stdin = cmd.InOrStdin()
		chat.Stdout = cmd.OutOrStdout()
		chat.Stderr = cmd.ErrOrStderr()
		err := chat.Run()
		if err != nil {
			cmd.PrintErrf("error running chat: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().StringVarP(&chatVersion, "version", "v", "cv1", "specify chat version to run (cv1, sv1, pv1)")
}
