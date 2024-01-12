package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(cliCmd)
}

var (
	cliCmd = &cobra.Command{
		Use:   "cli",
		Short: "run workflow in cli mode",
		Long:  `run workflow in cli mode`,
		Run: func(cmd *cobra.Command, args []string) {
			// run cli loop
			cli := exec.Command("go", "run", "cmd/cli/main.go")
			// pass stdin, stdout, and stderr to child process
			cli.Stdin = os.Stdin
			cli.Stdout = os.Stdout
			cli.Stderr = os.Stderr
			err := cli.Run()
			if err != nil {
				log.Fatalf("error running cli: %s", err)
			}
			fmt.Println("cli exited")
		},
	}
)
