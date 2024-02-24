package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/logrusorgru/aurora"
	util "github.com/ooyeku/flow/internal/util"
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

// Run starts the non-streaming chat application loop.
// It displays a welcome message and a stylish box.
// It then enters a loop where it prompts the user for input, generates an AI response using the model,
// and saves the user input and AI response to the chat store.
// The loop continues until the user enters "exit" or encounters an error.
// It returns nil if the user enters "exit" or encounters no errors.
// Otherwise, it returns an error indicating the cause.
func (app *ChatApp) Run() error {
	au := aurora.NewAurora(true)

	// Welcome message with stylish box
	fmt.Println(au.Bold(au.BgCyan(" Welcome to Go Flow Chat (Non-Streaming) ")))
	fmt.Println(au.Bold(au.BgCyan("---------------------------------------")))

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
		aiResponse = util.RemoveCurlyBraces(aiResponse)
		fmt.Print(aiResponse)

		// Save chat history
		err = app.chatStore.SaveEntry(userInput, aiResponse)
		if err != nil {
			return fmt.Errorf("error saving chat entry: %w", err)
		}
	}
}

// Close gracefully shuts down the chat application. It closes the genai client and chat store connections.
func (app *ChatApp) Close() {
	if err := app.client.Close(); err != nil {
		log.Printf("Error closing genai client: %s", err)
	}
	if err := app.chatStore.Close(); err != nil {
		log.Printf("Error closing chat store: %s", err)
	}
}

// main is the entry point function for the chat application.
// It loads the configuration, creates a chat app, sets up signal handling, and starts the chat loop.
// It waits for a shutdown signal before exiting.
//
// The chat loop runs in a separate goroutine and continuously prompts the user for input.
// It sends the user input to the AI model to generate a response, and then displays the response.
// The chat loop also saves the user input and AI response to a chat store for future retrieval.
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
