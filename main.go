package main

import (
	"github.com/ooyeku/flow/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
