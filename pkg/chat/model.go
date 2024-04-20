package chat

import (
	"bufio"
	"github.com/google/uuid"
	"net/http"
	_ "os"
	"time"
)

type Thread struct {
	ID      uuid.UUID
	Created time.Time
	Name    string
	Msgs    []ChatResponse
}

func NewThread(name string) *Thread {
	return &Thread{
		ID:      uuid.New(),
		Created: time.Now(),
		Name:    name,
	}
}

type ChatTopic struct {
	ID          uuid.UUID `'json:"id"`
	Name        string    `'json:"name"`
	Description string    `'json:"description"`
	Threads     []Thread  `'json:"threads"`
}

func NewChatTopic(name, description string) *ChatTopic {
	return &ChatTopic{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}
}

type ChatApp struct {
	Client       *http.Client
	ChatService  *ChatService
	Scanner      *bufio.Scanner
	ApiKey       string
	CurrentModel string
	Models       []string
	CurrentTopic ChatTopic
	Topics       []*ChatTopic
}

func NewChatApp(client *http.Client, chatService *ChatService, scanner *bufio.Scanner, apiKey string, models []string) (*ChatApp, error) {
	return &ChatApp{
		Client:       client,
		ChatService:  chatService,
		Scanner:      scanner,
		ApiKey:       apiKey,
		CurrentModel: models[0],
		Models:       models,
	}, nil
}

type Primer struct {
	System []map[string]string
	User   []map[string]string
}

// Message is the message to be sent to the chatbot
type Message struct {
	ID         uuid.UUID
	Model      string
	Primer     Primer
	Temprature float64
	TopP       float64
	FreqPen    float64
	Stream     bool
	Query      string
}

func NewMessage(query string, model string) Message {
	return Message{
		ID:    uuid.New(),
		Model: model,
		Query: query,
	}
}

// ChatResponse to be received and stored
type ChatResponse struct {
	ID        string     `json:"id"`
	Model     string     `json:"model"`
	Time      string     `json:"time"`
	UserQuery Message    `json:"user_query"`
	Object    string     `json:"object"`
	Topic     *ChatTopic `json:"topic"`
}
