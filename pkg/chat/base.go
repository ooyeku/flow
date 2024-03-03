package chat

import (
	"github.com/asdine/storm"
	_ "github.com/sashabaranov/go-openai"
	"time"
)

type Entry struct {
	ID         int       `json:"id" storm:"id,increment"`
	Timestamp  time.Time `json:"timestamp"`
	UserInput  string    `json:"user_input"`
	AIResponse Response  `json:"ai_response"`
}

type Response struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Created string `json:"created"`
	Object  string `json:"object"`

	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

type OResponse struct {
	ID    string `json:"id"`
	Model string `json:"model"`
	Time  string `json:"time"`
	Text  string `json:"text"`
}

type ChatStore struct {
	db *storm.DB
}

func NewChatStore(dbPath string) (*ChatStore, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	return &ChatStore{db}, nil
}

func (cs *ChatStore) Close() error {
	return cs.db.Close()
}

func (cs *ChatStore) SaveEntry(userInput string, aiResponse Response) error {
	entry := &Entry{
		Timestamp:  time.Now(),
		UserInput:  userInput,
		AIResponse: aiResponse,
	}
	return cs.db.Save(entry)
}

func (cs *ChatStore) RetrieveEntries() ([]Entry, error) {
	var entries []Entry
	err := cs.db.All(&entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (cs *ChatStore) ClearEntries() (err error) {
	err = cs.db.Drop(&Entry{})
	return err
}
