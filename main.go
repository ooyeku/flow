package main

import (
	"flow/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
