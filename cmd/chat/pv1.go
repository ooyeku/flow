package main

import (
	"bufio"
	"bytes"
	_ "context"
	"encoding/json"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/ooyeku/flow/pkg/chat"
	"io"
	_ "io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type ChatAppP struct {
	client    *http.Client
	chatStore *chat.ChatStore
	scanner   *bufio.Scanner
	apikey    string
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
	}, nil
}

func (app *ChatAppP) RunP() error {
	au := aurora.NewAurora(true)

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

		payload := map[string]interface{}{
			"model": "pplx-70b-online",
			"messages": []map[string]string{
				{"role": "system", "content": "Be precise and concise."},
				{"role": "user", "content": userInput},
			},
		}
		payloadBytes, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", "https://api.perplexity.ai/chat/completions", bytes.NewReader(payloadBytes))
		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "Bearer "+app.apikey)

		res, _ := app.client.Do(req)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Printf("Error closing response body: %s", err)
			}
		}(res.Body)

		body, _ := io.ReadAll(res.Body)
		var result map[string]interface{}
		_ = json.Unmarshal(body, &result)

		aiResponse := fmt.Sprintf("AI: %s\n", au.Blue(body))
		fmt.Println(aiResponse)

		err := app.chatStore.SaveEntry(userInput, aiResponse)
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
