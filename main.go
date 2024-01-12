package main

import (
	"goworkflow/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
