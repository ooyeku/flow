package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/logrusorgru/aurora"
	util "github.com/ooyeku/flow/internal/util"
	"github.com/ooyeku/flow/pkg/chat"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type ChatConfigS struct {
	ModelName string
	ApiKey    string
	DbPath    string
}

type ChatAppS struct {
	config    *ChatConfigS
	client    *genai.Client
	model     *genai.GenerativeModel
	chatStore *chat.ChatStore
	scanner   *bufio.Scanner
}

func NewChatAppS(config *ChatConfigS) (*ChatAppS, error) {
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

	return &ChatAppS{
		config:    config,
		client:    client,
		model:     model,
		chatStore: chatStore,
		scanner:   bufio.NewScanner(os.Stdin),
	}, nil
}

func (app *ChatAppS) RunS() error {
	au := aurora.NewAurora(true)

	fmt.Println(au.Bold(au.BgCyan(" Welcome to Go Flow Chat (Streaming) ")))
	fmt.Println(au.Bold(au.BgCyan("-------------------------------------")))

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
			default:
				fmt.Println("Unknown command")
			}
			continue
		}

		// Rest of the chat loop...
		stream := app.model.GenerateContentStream(context.Background(), genai.Text(userInput))

		fmt.Print("AI: ")
		for {
			response, err := stream.Next()
			if errors.Is(err, iterator.Done) {
				fmt.Println()
				break
			}
			if err != nil {
				return fmt.Errorf("error getting next response: %w", err)
			}

			fmt.Print(au.Blue("."))
			strRes := util.RemoveCurlyBraces(fmt.Sprintf("AI: %s", response.Candidates[0].Content))
			fmt.Printf("%s", response.Candidates[0].Content)

			err = app.chatStore.SaveEntry(userInput, strRes)
			if err != nil {
				return fmt.Errorf("error saving chat entry: %w", err)
			}
		}
	}
}

func (app *ChatAppS) CloseS() {
	if err := app.client.Close(); err != nil {
		log.Printf("Error closing genai client: %s", err)
	}
	if err := app.chatStore.Close(); err != nil {
		log.Printf("Error closing chat store: %s", err)
	}
}

func main() {
	config := &ChatConfigS{
		ModelName: "gemini-1.0-pro",
		ApiKey:    os.Getenv("GOOGLE_AI_STUDIO"),
		DbPath:    "internal/inmemory/sv1.db",
	}

	app, err := NewChatAppS(config)
	if err != nil {
		log.Fatalf("Error creating chat app: %s", err)
	}
	defer app.CloseS()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.RunS(); err != nil {
			log.Println("Chat loop error:", err)
		}
		signals <- syscall.SIGINT
	}()

	<-signals
}
