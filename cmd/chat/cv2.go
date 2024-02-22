package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup signal handling to gracefully shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Start the chat loop in a separate goroutine, so we can listen for the interrupt signal concurrently.
	go func() {
		if err := chatLoop2(ctx); err != nil {
			log.Println("Chat loop error:", err)
			cancel()
		}
	}()

	select {
	case <-signals:
		color.Yellow("Shutting down gracefully...")
	case <-ctx.Done():
		// It's possible the context was already cancelled due to an error in the chat loop.
	}

	color.Green("Goodbye!")
}

func chatLoop2(ctx context.Context) error {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_AI_STUDIO")))
	if err != nil {
		return fmt.Errorf("failed to create a client: %w", err)
	}
	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			_ = fmt.Errorf("error closing genai client: %w", err)
		}
	}(client)

	model := client.GenerativeModel("gemini-1.0-pro")

	scanner := bufio.NewScanner(os.Stdin)

	color.Cyan("Welcome to the chat! Type your message and press enter to chat with the AI. Type 'exit' to quit.")

	for {
		color.Green("You: ")
		scanned := scanner.Scan()
		if !scanned {
			return scanner.Err()
		}

		userInput := scanner.Text()

		if userInput == "exit" {
			return nil
		}

		response, err := model.GenerateContent(ctx, genai.Text(userInput))
		if err != nil {
			return fmt.Errorf("error generating response: %w", err)
		}

		aiResponse := fmt.Sprintf("AI: %s\n", response.Candidates[0].Content)
		color.Blue(aiResponse)

		select {
		case <-ctx.Done():
			return nil
		default:
			// Continue the loop if context is not done.
		}
	}
}
