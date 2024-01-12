package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Welcome to the workflow CLI!")
		fmt.Println("Please enter a command (or 'exit' to quit):")

		// Read the user's input
		input, _ := reader.ReadString('\n')

		// Remove trailing newline
		input = strings.TrimSuffix(input, "\n")

		// Handle the input
		if input == "exit" {
			break
		} else {
			fmt.Printf("You entered: %s\n", input)
		}
	}
}
