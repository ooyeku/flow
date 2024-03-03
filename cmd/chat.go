package cmd

import (
	"github.com/spf13/cobra"
	"os/exec"
)

var chatVersion string

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Launch a chat with the AI",
	Long: `Launch a chat with the AI using the pv1(perplexity).

chat history will be stored in a local chat.db file and can be accessed by running scripts/getchathistory.go

The chat will run until you type 'exit' and press enter.

Example usage:
go run main.go chat --version pv1
go run main.go chat -v pv1`,
	Run: func(cmd *cobra.Command, args []string) {
		var chat *exec.Cmd
		if chatVersion == "pv1" {
			chat = exec.Command("bash", "-c", "./pv1")
		} else if chatVersion == "pv2" {
			chat = exec.Command("bash", "-c", "./pv2")
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
	chatCmd.Flags().StringVarP(&chatVersion, "version", "v", "pv1", "specify chat version to run (pv1)")
}
