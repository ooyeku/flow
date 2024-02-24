package main

import (
	"bufio"
	_ "bufio"
	"context"
	"errors"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/logrusorgru/aurora"
	"github.com/ooyeku/flow/pkg/chat"
	_ "github.com/ooyeku/flow/pkg/chat"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	_ "google.golang.org/api/option"
	_ "io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ChatConfigS holds configuration options for the chat application.
type ChatConfigS struct {
	ModelName string
	ApiKey    string
	DbPath    string
}

// ChatAppS represents the chat application with streaming response handling.
type ChatAppS struct {
	config    *ChatConfigS
	client    *genai.Client
	model     *genai.GenerativeModel
	chatStore *chat.ChatStore
	scanner   *bufio.Scanner
}

func NewChatAppS(config *ChatConfigS) (*ChatAppS, error) {
	chatStore, err := chat.NewChatStore(config.DbPath)
	if err != nil {
		return nil, fmt.Errorf("error creating chat store: %w", err)
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.ApiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create a client: %w", err)
	}

	model := client.GenerativeModel(config.ModelName)

	scanner := bufio.NewScanner(os.Stdin)

	return &ChatAppS{
		config:    config,
		client:    client,
		model:     model,
		chatStore: chatStore,
		scanner:   scanner,
	}, nil
}

// Run starts the chat application loop with streaming response handling.
func (app *ChatAppS) RunS() error {
	au := aurora.NewAurora(true)

	fmt.Println(au.Cyan("Welcome to the chat! Type your message and press enter to chat with the AI. Type 'exit' to quit."))

	for {
		fmt.Print(au.Green("You: "))
		scanned := app.scanner.Scan()
		if !scanned {
			return app.scanner.Err()
		}

		userInput := app.scanner.Text()

		if userInput == "exit" {
			return nil
		}

		// Handle multi-line input or commands here...

		// Generate response stream
		stream := app.model.GenerateContentStream(context.Background(), genai.Text(userInput))

		// Handle streaming response
		for {
			response, err := stream.Next()
			if errors.Is(err, iterator.Done) {
				break
			}
			if err != nil {
				return fmt.Errorf("error getting next response: %w", err)
			}

			aiResponse := fmt.Sprintf("AI: %s", au.Blue(response.Candidates[0].Content))
			fmt.Println(aiResponse)

			// Save chat history
			err = app.chatStore.SaveEntry(userInput, aiResponse)
			if err != nil {
				return fmt.Errorf("error saving chat entry: %w", err)
			}
		}
	}
}

// ... (Close and main functions remain the same)
func (app *ChatAppS) CloseS() {
	if err := app.client.Close(); err != nil {
		log.Printf("Error closing genai client: %s", err)
	}
	if err := app.chatStore.Close(); err != nil {
		log.Printf("Error closing chat store: %s", err)
	}
}

func main() {
	// Load configuration from file or environment variables...
	config := &ChatConfigS{
		ModelName: "gemini-1.0-pro",
		ApiKey:    os.Getenv("GOOGLE_AI_STUDIO"),
		DbPath:    "chat.db",
	}

	app, err := NewChatAppS(config)
	if err != nil {
		log.Fatalf("Error creating chat app: %s", err)
	}
	defer app.CloseS()

	// Setup signal handling
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.RunS(); err != nil {
			log.Println("Chat loop error:", err)
		}
		signals <- syscall.SIGINT // Signal shutdown after chat loop exits
	}()

	<-signals // Wait for shutdown signal
}
