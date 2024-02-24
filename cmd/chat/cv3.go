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

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := chatLoop3(ctx); err != nil {
			log.Println("Chat loop error:", err)
			cancel()
		}
	}()

	select {
	case <-signals:
		color.Yellow("Shutting down gracefully...")
	case <-ctx.Done():
	}

	color.Green("Goodbye!")
}

func chatLoop3(ctx context.Context) error {
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

	model := client.GenerativeModel("gemma")

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

		responseStream := make(chan *genai.GenerateContentResponse, 1)
		errorStream := make(chan error, 1)

		go func() {
			response, err := model.GenerateContent(ctx, genai.Text(userInput)) // hypothetical method
			if err != nil {
				errorStream <- fmt.Errorf("error generating response: %w", err)
				return
			}
			responseStream <- response
			close(responseStream)
		}()

		select {
		case response := <-responseStream:
			aiResponse := fmt.Sprintf("AI: %s\n", response.Candidates[0].Content)
			color.Blue(aiResponse)
		case err := <-errorStream:
			return err
		case <-ctx.Done():
			return nil
		}
	}
}
