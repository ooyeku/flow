package chat

import (
	"github.com/asdine/storm"
	"time"
)

type Entry struct {
	ID         int       `json:"id" storm:"id,increment"`
	Timestamp  time.Time `json:"timestamp"`
	UserInput  string    `json:"user_input"`
	AIResponse string    `json:"ai_response"`
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

func (cs *ChatStore) SaveEntry(userInput, aiResponse string) error {
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
