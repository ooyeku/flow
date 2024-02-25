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
	"syscall"
)

// ChatConfigS holds configuration options for the chat application.
type ChatConfigS struct {
	ModelName string
	ApiKey    string
	DbPath    string
}

// ChatAppS represents the chat application.
type ChatAppS struct {
	config    *ChatConfigS
	client    *genai.Client
	model     *genai.GenerativeModel
	chatStore *chat.ChatStore
	scanner   *bufio.Scanner
}

// NewChatAppS is a function that creates a new instance of ChatAppS.
// It takes a ChatConfigS pointer as a parameter, which holds the configuration options for the chat application.
// The function returns a pointer to ChatAppS and an error.
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

// RunS starts the streaming chat application loop.
// It prints a welcome message with a stylish box and prompts the user for input.
// If the input is "exit", it returns nil to exit the loop.
// Otherwise, it generates a response stream based on the user input using the generative model.
// While receiving responses from the stream, it prints animated dots and saves the chat history.
// After the response is complete, it adds a newline and continues the loop.
// If any error occurs during the process, it returns an error.
func (app *ChatAppS) RunS() error {
	au := aurora.NewAurora(true)

	// Welcome message with stylish box
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

		// Handle multi-line input or commands here...

		// Generate response stream
		stream := app.model.GenerateContentStream(context.Background(), genai.Text(userInput))

		// Handle streaming response with animated dots
		fmt.Print("AI: ")
		for {
			response, err := stream.Next()
			if errors.Is(err, iterator.Done) {
				fmt.Println() // Add newline after response is complete
				break
			}
			if err != nil {
				return fmt.Errorf("error getting next response: %w", err)
			}
			fmt.Print(au.Blue(".")) // Print animated dots while receiving
			// Save chat history
			// change response.Candidates[0].Content to string
			strRes := util.RemoveCurlyBraces(fmt.Sprintf("AI: %s", response.Candidates[0].Content))
			fmt.Printf("%s", response.Candidates[0].Content)

			err = app.chatStore.SaveEntry(userInput, strRes)
			if err != nil {
				return fmt.Errorf("error saving chat entry: %w", err)
			}
		}
	}
}

// CloseS closes the chat application by closing the genai client and the chat store. It logs any errors that occur during the closing process.
// Usage example:
// ```
// // ... (Close and main functions remain the same)
//
//	func main() {
//		// Load configuration from file or environment variables...
//		config := &ChatConfigS{
//			ModelName: "gemini-1.0-pro",
//			ApiKey:    os.Getenv("GOOGLE_AI_STUDIO"),
//			DbPath:    "chat.db",
//		}
//
//		app, err := NewChatAppS(config)
//		if err != nil {
//			log.Fatalf("Error creating chat app: %s", err)
//		}
//		defer app.CloseS()
//
//		// Setup signal handling
//		signals := make(chan os.Signal, 1)
//		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
//
//		go func() {
//			if err := app.RunS(); err != nil {
//				log.Println("Chat loop error:", err)
//			}
//			signals <- syscall.SIGINT // Signal shutdown after chat loop exits
//		}()
//
//		<-signals // Wait for shutdown signal
//	}
//
// ```
func (app *ChatAppS) CloseS() {
	if err := app.client.Close(); err != nil {
		log.Printf("Error closing genai client: %s", err)
	}
	if err := app.chatStore.Close(); err != nil {
		log.Printf("Error closing chat store: %s", err)
	}
}

// main is the entry point of the chat application.
// It loads the configuration, creates a new ChatAppS instance, sets up signal handling, and starts the chat loop.
// The chat loop runs in a separate goroutine and waits for a shutdown signal from the main goroutine.
// Once the shutdown signal is received, the chat loop exits and the application closes.
func main() {
	// Load configuration from file or environment variables...
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
