package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/ooyeku/flow/pkg/chat"
	"github.com/sashabaranov/go-openai"
	"github.com/theckman/yacspin"
	"log"
	"os"
	"strings"
	"time"
)

type ChatAppO struct {
	client    *openai.Client
	chatStore *chat.ChatStore
	scanner   *bufio.Scanner
	apikey    string
	model     string
}

func NewChatAppO(dbPath string) (*ChatAppO, error) {
	chatStore, err := chat.NewChatStore(dbPath)
	if err != nil {
		return nil, err
	}

	apikey := os.Getenv("OPENAI_API_KEY")
	if apikey == "" {
		return nil, fmt.Errorf("OPENAI_KEY environment variable not set")
	}

	client := openai.NewClient(apikey)

	return &ChatAppO{
		client:    client,
		chatStore: chatStore,
		scanner:   bufio.NewScanner(os.Stdin),
		apikey:    apikey,
		model:     openai.GPT4Turbo0125, // default model
	}, nil
}

func (app *ChatAppO) RunP() error {
	au := aurora.NewAurora(true)

	// Similar logic to pv2.go's RunP function
	// ...
	cfg := yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[11],
		Suffix:          au.Bold(au.BrightBlue("Pondering...")).String(),
		SuffixAutoColon: true,
		StopCharacter:   "💡",
	}

	spinner, err := yacspin.New(cfg)
	if err != nil {
		log.Fatal("error creating spinner: ", err)
	}

	intro()
	for {
		fmt.Println(au.Bold(au.Gray(12, "Current model: ")), au.Bold(au.BrightGreen(app.model)))
		fmt.Print(au.Green("You: "))
		scanned := app.scanner.Scan()
		if !scanned {
			return app.scanner.Err()
		}

		prompt := app.scanner.Text()
		if prompt == "exit" {
			fmt.Println(au.Bold(au.BrightBlue("Goodbye!")))
			return nil
		}

		if strings.HasPrefix(prompt, "$") {
			switch prompt {
			case "$history":
				history, err := app.chatStore.RetrieveEntries()
				if err != nil {
					return err
				}
				// handle empty history
				if len(history) == 0 {
					fmt.Println(au.Bold(au.BrightBlue("No chat history")))
					continue
				}
				for _, entry := range history {
					fmt.Println(au.Bold(au.BrightGreen("You: ")), au.Green(entry.UserInput))
					fmt.Println(au.Bold(au.BrightBlue("AI: ")), entry.AIResponse.Choices[0].Message.Content)
					fmt.Println(au.Bold(au.BrightBlue("Time: ")), entry.Timestamp)
					fmt.Println(au.Bold(au.BrightBlue("Model: ")), app.model)
					// print a line to separate chat history entries
					fmt.Println(au.Bold(au.BrightBlue("-------------------------------------------------")))

				}
				continue
			default:
				fmt.Println(au.Bold(au.BrightRed("Invalid command")))
			}
			continue
		}

		spinner.Start()
		resp, err := app.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model: app.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
		)
		if err != nil {
			return err
		}
		spinner.Stop()
		aiResponse := resp.Choices[0].Message.Content
		fmt.Println(au.Bold(au.BrightBlue("AI: ")), aiResponse)

		// save chat history
		var result chat.Response
		result.ID = resp.ID
		result.Model = app.model
		result.Created = time.Now().Format(time.RFC3339)

		err = app.chatStore.SaveEntry(prompt, result)
		if err != nil {
			return err
		}

	}

	return nil
}

func (app *ChatAppO) CloseP() {
	if err := app.chatStore.Close(); err != nil {
		log.Printf("Error closing chat store: %s", err)
	}
}

func main() {
	dbPath := "internal/inmemory/ov1.db"

	app, err := NewChatAppO(dbPath)
	if err != nil {
		log.Fatalf("Error creating chat app: %s", err)
	}
	defer app.CloseP()

	if err := app.RunP(); err != nil {
		log.Fatalf("Error running chat app: %s", err)
	}
}

func intro() {
	au := aurora.NewAurora(true)
	fmt.Println(au.Bold(au.Cyan(" Welcome to Flow Chat (OpenAI) ")))
	fmt.Println(au.Bold(au.BgCyan("-----------------------------------------")))
	// options
	fmt.Println(au.Gray(12, "Type 'exit' or $exit to quit the chat"))
	fmt.Println(au.Gray(12, "Type $history to view chat history"))
}
