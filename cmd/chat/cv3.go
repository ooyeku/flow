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

// ChatConfig holds configuration options for the chat application.
type ChatConfig struct {
	ModelName string
	ApiKey    string
	DbPath    string
}

// ChatApp represents the chat application.
type ChatApp struct {
	config    *ChatConfig
	client    *genai.Client
	model     *genai.GenerativeModel
	chatStore *chat.ChatStore
	scanner   *bufio.Scanner
}

// NewChatApp creates a new instance of ChatApp with the provided configuration.
func NewChatApp(config *ChatConfig) (*ChatApp, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(config.ApiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	model := client.GenerativeModel(config.ModelName)

	chatStore, err := chat.NewChatStore(config.DbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat store: %w", err)
	}

	return &ChatApp{
		config:    config,
		client:    client,
		model:     model,
		chatStore: chatStore,
		scanner:   bufio.NewScanner(os.Stdin),
	}, nil
}

// Run starts the chat application loop.
func (app *ChatApp) Run() error {
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

		response, err := app.model.GenerateContent(context.Background(), genai.Text(userInput))
		if err != nil {
			return fmt.Errorf("error generating response: %w", err)
		}

		aiResponse := fmt.Sprintf("AI: %s\n", au.Blue(response.Candidates[0].Content))
		fmt.Print(aiResponse)

		// Save chat history
		err = app.chatStore.SaveEntry(userInput, aiResponse)
		if err != nil {
			return fmt.Errorf("error saving chat entry: %w", err)
		}
	}
}

// Close gracefully shuts down the chat application.
func (app *ChatApp) Close() {
	if err := app.client.Close(); err != nil {
		log.Printf("Error closing genai client: %s", err)
	}
	if err := app.chatStore.Close(); err != nil {
		log.Printf("Error closing chat store: %s", err)
	}
}

func main() {
	// Load configuration from file or environment variables...
	config := &ChatConfig{
		ModelName: "gemini-1.0-pro",
		ApiKey:    os.Getenv("GOOGLE_AI_STUDIO"),
		DbPath:    "chat.db",
	}

	app, err := NewChatApp(config)
	if err != nil {
		log.Fatalf("Error creating chat app: %s", err)
	}
	defer app.Close()

	// Setup signal handling
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Run(); err != nil {
			log.Println("Chat loop error:", err)
		}
		signals <- syscall.SIGINT // Signal shutdown after chat loop exits
	}()

	<-signals // Wait for shutdown signal
}
