package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/logrusorgru/aurora"
	"github.com/ooyeku/flow/pkg/chat"
	"google.golang.org/api/option"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var au = aurora.NewAurora(true) // Assuming you want to enable colors, otherwise pass false

// var commandEnterPressed bool

func main() {
	chatStore, err := chat.NewChatStore("chat.db")
	if err != nil {
		log.Fatalf("error creating chat store: %s", err)
	}
	defer func(chatStore *chat.ChatStore) {
		err := chatStore.Close()
		if err != nil {
			log.Fatalf("error closing chat store: %s", err)
		}
	}(chatStore)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup signal handling to gracefully shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Start the chat loop in a separate goroutine, so we can listen for the interrupt signal concurrently.
	go func() {
		if err := chatLoop(ctx, chatStore); err != nil {
			log.Println("Chat loop error:", err)
			cancel()
		}
	}()

	select {
	case <-signals:
		fmt.Println("Shutting down gracefully...")
	case <-ctx.Done():
		// It's possible the context was already cancelled due to an error in the chat loop.
	}

	fmt.Println("Goodbye!")
}

func chatLoop(ctx context.Context, chatStore *chat.ChatStore) error {
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

	for {
		fmt.Print(au.Colorize("You: ", aurora.GreenFg))
		scanned := scanner.Scan()
		if !scanned {
			return scanner.Err()
		}

		userInput := scanner.Text()

		// // Check if userInput is empty, indicating the user only pressed enter
		// if userInput == "" {
		// 	if !commandEnterPressed {
		// 		commandEnterPressed = true
		// 		continue
		// 	} else {
		// 		return nil
		// 	}
		// }
		// commandEnterPressed = false

		response, err := model.GenerateContent(ctx, genai.Text(userInput))
		if err != nil {
			return fmt.Errorf("error generating response: %w", err)
		}

		aiResponse := fmt.Sprintf("AI: %s\n", au.Colorize(response.Candidates[0].Content, aurora.BlueFg))
		fmt.Print(aiResponse)

		// convert response.Candidates[0].Content to string
		aiReHis := fmt.Sprintf("AI: %s\n", response.Candidates[0].Content)

		err = chatStore.SaveEntry(userInput, aiReHis)
		if err != nil {
			return fmt.Errorf("error saving chat entry: %w", err)
		}

		select {
		case <-ctx.Done():
			return nil
		default:
			// Continue the loop if context is not done.
		}
	}
}
