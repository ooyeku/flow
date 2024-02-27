package main

import (
	"bufio"
	"bytes"
	_ "context"
	"encoding/json"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/ooyeku/flow/pkg/chat"
	"github.com/theckman/yacspin"
	"io"
	_ "io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// model options
const sonarSmallChat = "sonar-small-chat"
const sonarSmallOnline = "sonar-small-online"
const sonarMediumChat = "sonar-medium-chat"
const sonarMediumOnline = "sonar-medium-online"
const pplx70b = "pplx-70b-online"
const codeLLama70b = "code-llama-70b-instruct"
const mixtral7b = "mixtral-7b-instruct"
const mixtral8x7b = "mixtral-8x7b-instruct"

type ChatAppP struct {
	client    *http.Client
	chatStore *chat.ChatStore
	scanner   *bufio.Scanner
	apikey    string
	model     string
}

func NewChatAppP(dbPath string) (*ChatAppP, error) {
	chatStore, err := chat.NewChatStore(dbPath)
	if err != nil {
		return nil, err
	}

	// Get the API key from the environment
	apikey := os.Getenv("PAI_KEY")
	if apikey == "" {
		return nil, fmt.Errorf("PAI_KEY environment variable not set")
	}

	return &ChatAppP{
		client:    &http.Client{},
		chatStore: chatStore,
		scanner:   bufio.NewScanner(os.Stdin),
		apikey:    apikey,
		model:     pplx70b, // default model
	}, nil
}

func (app *ChatAppP) RunP() error {
	au := aurora.NewAurora(true)

	// spinner configuration
	cfg := yacspin.Config{
		Frequency:       100 * time.Microsecond,
		CharSet:         yacspin.CharSets[59],
		Suffix:          au.Bold(au.BrightBlue("Pondering...")).String(),
		SuffixAutoColon: true,
		StopCharacter:   "💡",
	}

	spinner, err := yacspin.New(cfg)
	if err != nil {
		log.Fatal("error creating spinner: ", err)
	}

	fmt.Println(au.Bold(au.BgCyan(" Welcome to Go Flow Chat (Perplexity AI) ")))
	fmt.Println(au.Bold(au.BgCyan("-----------------------------------------")))

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
				fmt.Print(au.Bold(au.Red("Are you sure you want to clear chat history? (yes/no): ")))
				scanned := app.scanner.Scan()
				if !scanned {
					return app.scanner.Err()
				}
				confirmation := app.scanner.Text()
				if confirmation != "yes" {
					fmt.Println(au.Red("Chat history not cleared"))
					continue
				} else {
					err := app.chatStore.ClearEntries()
					if err != nil {
						return fmt.Errorf("error clearing chat history: %w", err)
					}
					fmt.Println(au.Bold(au.Green("Chat history cleared")))
				}
			case "model":
				fmt.Print(au.Bold(au.BrightBlue("Enter model name: ")))
				// display model options
				fmt.Println(au.Bold(au.BrightBlue("Model options: ")))
				fmt.Println(au.Bold(au.BrightGreen("sonar-small-chat")))
				fmt.Println(au.Bold(au.BrightGreen("sonar-small-online")))
				fmt.Println(au.Bold(au.BrightGreen("sonar-medium-chat")))
				fmt.Println(au.Bold(au.BrightGreen("sonar-medium-online")))
				fmt.Println(au.Bold(au.BrightGreen("pplx-70b")))
				fmt.Println(au.Bold(au.BrightGreen("code-llama-70b")))
				fmt.Println(au.Bold(au.BrightGreen("mixtral-7b")))
				fmt.Println(au.Bold(au.BrightGreen("mixtral-8x7b")))
				scanned := app.scanner.Scan()
				if !scanned {
					return app.scanner.Err()
				}

				app.model = app.scanner.Text()
				fmt.Println(au.Bold(au.BrightBlue("Model set to: ")), au.Bold(au.BrightGreen(app.model)))
				continue
			default:
				fmt.Println("Unknown command")
			}
			continue
		}

		// Rest of the chat loop...
		payload := map[string]interface{}{
			"model": app.model,
			"messages": []map[string]string{
				{"role": "system", "content": "Be precise and concise."},
				{"role": "user", "content": userInput},
			},
			"stream": "false",
		}
		payloadBytes, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", "https://api.perplexity.ai/chat/completions", bytes.NewReader(payloadBytes))
		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "Bearer "+app.apikey)

		err := spinner.Start()
		if err != nil {
			log.Fatal("error starting spinner: ", err)
		}

		res, _ := app.client.Do(req)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Printf("Error closing response body: %s", err)
			}
		}(res.Body)

		body, _ := io.ReadAll(res.Body)
		var result chat.Response
		_ = json.Unmarshal(body, &result)

		err = spinner.Stop()
		if err != nil {
			log.Fatal("error stopping spinner: ", err)
		}

		ai := au.Bold(au.BrightBlue("AI: "))
		aiResponse := fmt.Sprintf("%s%s", ai, result.Choices[0].Message.Content)
		fmt.Println(aiResponse)

		err = app.chatStore.SaveEntry(userInput, result)
		if err != nil {
			return err
		}
	}
}

func (app *ChatAppP) CloseP() {
	if err := app.chatStore.Close(); err != nil {
		log.Printf("Error closing chat store: %s", err)
	}
}

func main() {
	dbPath := "internal/inmemory/pv1.db"

	app, err := NewChatAppP(dbPath)
	if err != nil {
		log.Fatalf("Error creating chat app: %s", err)
	}
	defer app.CloseP()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.RunP(); err != nil {
			log.Println("Chat loop error:", err)
		}
		signals <- syscall.SIGINT
	}()

	<-signals
}
