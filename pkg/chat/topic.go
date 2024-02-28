package chat

import (
	"github.com/asdine/storm"
	"time"
)

type Topic struct {
	ID   int    `json:"id" storm:"id,increment"`
	Name string `json:"name"`
}

func NewTopic(name string) *Topic {
	return &Topic{
		Name: name,
	}
}

type TEntry struct {
	ID         int       `json:"id" storm:"id,increment"`
	Timestamp  time.Time `json:"timestamp"`
	UserInput  string    `json:"user_input"`
	AIResponse TResponse `json:"ai_response"`
	Topic      Topic     `json:"topic"`
}

type TResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Created string `json:"created"`
	Object  string `json:"object"`

	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type TopicStore struct {
	db *storm.DB
}

func NewTopicStore(dbPath string) (*TopicStore, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	return &TopicStore{db}, nil
}

func (ts *TopicStore) Close() error {
	return ts.db.Close()
}

func (ts *TopicStore) SaveEntry(userInput string, aiResponse TResponse, topic Topic) error {
	entry := &TEntry{
		Timestamp:  time.Now(),
		UserInput:  userInput,
		AIResponse: aiResponse,
		Topic:      topic,
	}
	return ts.db.Save(entry)
}

func (ts *TopicStore) RetrieveEntries() ([]TEntry, error) {
	var entries []TEntry
	err := ts.db.All(&entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (ts *TopicStore) RetrieveEntriesByTopic(topic string) ([]TEntry, error) {
	var entries []TEntry
	err := ts.db.Find("Topic.Name", topic, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (ts *TopicStore) RetrieveTopics() ([]Topic, error) {
	var topics []Topic
	err := ts.db.All(&topics)
	if err != nil {
		return nil, err
	}

	return topics, nil
}

func (ts *TopicStore) SaveTopic(name string) error {
	topic := &Topic{
		Name: name,
	}
	return ts.db.Save(topic)
}

func (ts *TopicStore) RetrieveTopicByName(name string) (Topic, error) {
	var topic Topic
	err := ts.db.One("Name", name, &topic)
	if err != nil {
		return Topic{}, err
	}

	return topic, nil
}
