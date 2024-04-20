package main

import (
	"bufio"
	"bytes"
	_ "context"
	"encoding/json"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/ooyeku/flow/cmd/chat/helpers"
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

const sonarSmallChat = "sonar-small-chat"

const sonarSmallOnline = "sonar-small-online"

const sonarMediumChat = "sonar-medium-chat"

const sonarMediumOnline = "sonar-medium-online"

const pplx70b = "pplx-70b-online"

const codeLLama70b = "codellama-70b-instruct"

const mixtral7b = "mixtral-7b-instruct"

const mixtral8x7b = "mixtral-8x7b-instruct"

const mixtral8x22b = "mixtral-8x22b-instruct"

const llama370b = "llama-3-70b-instruct"

// ChatAppP ChatAppO is a type that represents a chat application. It has the following properties:
// - client: an HTTP client used for making API requests
// - chatStore: a chat store used for storing chat history
// - scanner: a scanner used for reading user input from the console
// - apikey: the API key used for authentication
// - model: the current model for generating AI responses
// It also has the following methods:
// - RunP: starts the chat application and handles user input and AI responses
// - CloseP: closes the chat store
// The ChatAppO type can be created using the NewChatAppO function, which takes a database path as a parameter and returns an instance of ChatAppO.
// The ChatStore type, which is used by ChatAppO, is a separate type that represents a chat store. It has methods for saving chat entries, retrieving chat history, and clearing chat history. It can be created using the NewChatStore function, which takes a database path as a parameter and returns an instance of ChatStore.
// The pplx70b constant represents the default model for the ChatAppO type.
// The helpers.Intro function prints an introduction message and provides instructions for using the chat application.
type ChatAppP struct {
	client    *http.Client
	chatStore *chat.ChatStore
	scanner   *bufio.Scanner
	apikey    string
	model     string
}

// NewChatAppP returns a new instance of ChatAppP with the specified database path.
// It initializes the chatStore using the NewChatStore function.
// It retrieves the API key from the environment using os.Getenv("PAI_KEY").
// If the API key is not set, it returns an error.
// It initializes and returns the ChatAppP struct with the client, chatStore, scanner, apikey, and model.
// The model is set to "pplx-70b-online" by default.
// If there is an error creating the chatStore, it returns the error.
//
// Example usage:
// app, err := NewChatAppP(dbPath)
//
//	if err != nil {
//	    log.Fatalf("Error creating chat app: %s", err)
//	}
//
// defer app.CloseP()
//
// signals := make(chan os.Signal, 1)
// signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
//
//	go func() {
//	    if err := app.RunP(); err != nil {
//	        log.Println("Chat loop error:", err)
//	    }
//	    signals <- syscall.SIGINT
//	}()
//
// <-signals
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

// RunP runs the chat application
func (app *ChatAppP) RunP() error {
	au := aurora.NewAurora(true)

	// spinner configuration
	cfg := yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[59],
		Suffix:          au.Bold(au.BrightBlue("Pondering...")).String(),
		SuffixAutoColon: true,
		StopCharacter:   "💡",
	}

	spinner, err := yacspin.New(cfg)
	if err != nil {
		log.Fatal("error creating spinner: ", err)
	}

	helpers.Intro()
	//// current model
	fmt.Println(au.Bold(au.BrightBlue("Model: ")), au.Bold(au.BrightGreen(app.model)))

	for {
		fmt.Print(au.Green("You: "))
		scanned := app.scanner.Scan()
		if !scanned {
			return app.scanner.Err()
		}

		userInput := app.scanner.Text()

		if userInput == "exit" {
			fmt.Println(au.Bold(au.BrightBlue("Goodbye!")))
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
					fmt.Println(au.Bold(au.BrightGreen("You: ")), au.Green(entry.UserInput))
					fmt.Println(au.Bold(au.BrightBlue("AI: ")), entry.AIResponse.Choices[0].Message.Content)
					// print model details
					fmt.Println(au.Bold(au.BrightYellow("Model: ")), au.Bold(au.Magenta(entry.AIResponse.Model)))
					// print separator
					fmt.Println(au.Bold(au.BrightBlue("-------------------------------------------------")))
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
			case "models":
				fmt.Println(au.Bold(au.BrightBlue("Enter model name: ")))
				// display model options
				helpers.DisplayModels()
				scanned := app.scanner.Scan()
				if !scanned {
					return app.scanner.Err()
				}

				app.model = app.scanner.Text()
				// check if model is valid
				switch app.model {
				case sonarSmallChat, sonarSmallOnline, sonarMediumChat, sonarMediumOnline, pplx70b, codeLLama70b, mixtral7b, mixtral8x7b, mixtral8x22b, llama370b:
				default:
					fmt.Println(au.Bold(au.BrightRed("Invalid model")))
					app.model = pplx70b // set back to default
					fmt.Println(au.Bold(au.BrightBlue("Model set to default: ")), au.Bold(au.BrightGreen(app.model)))
					continue
				}

				fmt.Println(au.Bold(au.BrightBlue("Model set to: ")), au.Bold(au.BrightGreen(app.model)))
				continue
			case "exit":
				// Goodbye message
				fmt.Println(au.Bold(au.BrightBlue("Goodbye!")))
				return nil
			default:
				fmt.Println("Unknown command")
			}
			continue
		}

		payload := map[string]interface{}{
			"model": app.model,
			"messages": []map[string]string{
				{"role": "system", "content": helpers.SysMessage},
				{"role": "user", "content": userInput},
			},
			"temperature":       helpers.Temperature,
			"top_p":             helpers.TopP,
			"frequency_penalty": helpers.FrequencyPenalty,
			"stream":            helpers.Stream,
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

// CloseP closes the chat store associated with the ChatAppP instance.
// The method first attempts to close the chat store by calling app.chatStore.Close().
// If an error occurs during the closing process, the error is logged using log.Printf.
//
// Example usage:
//
//	  dbPath := "internal/inmemory/pv1.db"
//		app, err := NewChatAppP(dbPath)
//		if err != nil {
//		    log.Fatalf("Error creating chat app: %s", err)
//		}
//		defer app.CloseP()
//
// Note: The CloseP method is defined on the ChatAppP struct. The ChatAppP struct has a field called chatStore
// which is an instance of the ChatStore struct. The ChatStore struct has a Close method that is invoked by app.chatStore.Close()
// in the CloseP method implementation.
func (app *ChatAppP) CloseP() {
	if err := app.chatStore.Close(); err != nil {
		log.Printf("Error closing chat store: %s", err)
	}
}

// main is the entry point of the program.
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
