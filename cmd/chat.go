package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var chatVersion string

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Launch a chat with the AI",
	Long: `Launch a chat with the AI using the simple_chat or cv2 version.


The simple_chat version uses the logrusorgru/aurora package to colorize the output.
The cv2 version uses the fatih/color package to colorize the output.

The chat will run until you type 'exit' and press enter.

Example usage:
go run main.go chat --version cv2 
go run main.go chat --v simple_chat`,
	Run: func(cmd *cobra.Command, args []string) {
		var chat *exec.Cmd
		if chatVersion == "cv2" {
			chat = exec.Command("go", "run", "cmd/chat/cv2.go")
		} else {
			chat = exec.Command("go", "run", "cmd/chat/simple_chat.go")
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
	chatCmd.Flags().StringVarP(&chatVersion, "version", "v", "simple_chat", "specify chat version to run (simple_chat or cv2)")
}
