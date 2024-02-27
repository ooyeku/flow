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
	"strings"
	"syscall"
)

type ChatConfig struct {
	ModelName string
	ApiKey    string
	DbPath    string
}

type ChatApp struct {
	config    *ChatConfig
	client    *genai.Client
	model     *genai.GenerativeModel
	chatStore *chat.ChatStore
	scanner   *bufio.Scanner
}

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

func (app *ChatApp) Run() error {
	au := aurora.NewAurora(true)

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

		// Check if the user input starts with $
		if strings.HasPrefix(userInput, "$") {
			// Handle command
			command := strings.TrimPrefix(userInput, "$")
			switch command {
			case "history":
				// Retrieve and display chat history
				entries, err := app.chatStore.RetrieveEntries()
				if err != nil {
					return err
				}
				for _, entry := range entries {
					fmt.Printf("You: %s\nAI: %s\n", entry.UserInput, entry.AIResponse)
				}
			case "clear":
				// Ask for confirmation
				fmt.Println(au.Bold(au.Red("Are you sure you want to clear chat history? (yes/no): ")))
				scanned := app.scanner.Scan()
				if !scanned {
					return app.scanner.Err()
				}
				confirmation := app.scanner.Text()
				if confirmation != "yes" {
					fmt.Println(au.Red("Chat history not cleared"))
					continue
				} else {
					// Clear chat history
					err := app.chatStore.ClearEntries()
					if err != nil {
						return fmt.Errorf("error clearing chat history: %w", err)
					}
					fmt.Println(au.Bold(au.Green("Chat history cleared")))
				}
			default:
				fmt.Println("Unknown command")
			}
			continue
		}

		response, err := app.model.GenerateContent(context.Background(), genai.Text(userInput))
		if err != nil {
			return fmt.Errorf("error generating response: %w", err)
		}

		aiResponse := fmt.Sprintf("AI: %s\n", au.Blue(response.Candidates[0].Content))
		aiResponse = util.RemoveCurlyBraces(aiResponse)
		fmt.Println(aiResponse)

		err = app.chatStore.SaveEntry(userInput, aiResponse)
		if err != nil {
			return fmt.Errorf("error saving chat entry: %w", err)
		}
	}
}

func (app *ChatApp) Close() {
	if err := app.client.Close(); err != nil {
		log.Printf("Error closing genai client: %s", err)
	}
	if err := app.chatStore.Close(); err != nil {
		log.Printf("Error closing chat store: %s", err)
	}
}

func main() {
	config := &ChatConfig{
		ModelName: "gemini-1.0-pro",
		ApiKey:    os.Getenv("GOOGLE_AI_STUDIO"),
		DbPath:    "internal/inmemory/cv1.db",
	}

	app, err := NewChatApp(config)
	if err != nil {
		log.Fatalf("Error creating chat app: %s", err)
	}
	defer app.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Run(); err != nil {
			log.Println("Chat loop error:", err)
		}
		signals <- syscall.SIGINT
	}()

	<-signals
}