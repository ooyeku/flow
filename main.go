package main

import (
	"github.com/ooyeku/flow/cmd"
	"os/exec"
)

func main() {
	exec.Command("go", "build", "cmd/cli/cliapp.go").Run()
	exec.Command("go", "build", "cmd/chat/pv1.go").Run()
	err := cmd.Execute()
	if err != nil {
		return
	}
}
