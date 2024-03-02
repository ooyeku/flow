package chat

import (
	"bufio"
	"github.com/google/uuid"
	"net/http"
)

type ChatService struct {
	chatStore PPLXChatStore
	ChatApp   *ChatApp
}

func NewChatService(cs PPLXChatStore) *ChatService {
	return &ChatService{chatStore: cs}
}

func (cs *ChatService) CreateThread(t *Thread) error {
	return cs.chatStore.CreateThread(t)
}

func (cs *ChatService) GetThread(ID uuid.UUID) (*Thread, error) {
	return cs.chatStore.GetThread(ID)
}

func (cs *ChatService) ListThreads() ([]*Thread, error) {
	return cs.chatStore.ListThreads()
}

func (cs *ChatService) UpdateThread(ID uuid.UUID, t *Thread) error {
	return cs.chatStore.UpdateThread(ID, t)
}

func (cs *ChatService) DeleteThread(ID uuid.UUID) error {
	return cs.chatStore.DeleteThread(ID)
}

func (cs *ChatService) CreateTopic(t *ChatTopic) error {
	return cs.chatStore.CreateTopic(t)
}

func (cs *ChatService) GetTopic(ID uuid.UUID) (*ChatTopic, error) {
	return cs.chatStore.GetTopic(ID)
}

func (cs *ChatService) ListTopics() ([]*ChatTopic, error) {
	return cs.chatStore.ListTopics()
}

func (cs *ChatService) DeleteTopic(ID uuid.UUID) error {
	return cs.chatStore.DeleteTopic(ID)
}

func (cs *ChatService) UpdateTopic(ID uuid.UUID, t *ChatTopic) error {
	return cs.chatStore.UpdateTopic(ID, t)
}

func (cs *ChatService) SaveChatResponse(cr *ChatResponse) error {
	return cs.chatStore.SaveChatResponse(cr)
}

func (cs *ChatService) GetChatResponse(ID string) (*ChatResponse, error) {
	return cs.chatStore.GetChatResponse(ID)
}

func (cs *ChatService) ListChatResponses() ([]*ChatResponse, error) {
	return cs.chatStore.ListChatResponses()
}

func (cs *ChatService) NewChatApp(client *http.Client, chatService *ChatService, scanner *bufio.Scanner, apiKey string, models []string) (*ChatApp, error) {
	return NewChatApp(client, chatService, scanner, apiKey, models)
}

func (cs *ChatService) ClearEntries() error {
	return cs.chatStore.ClearEntries()

}

func (cs *ChatService) ClearEntriesByTopic(name string) error {
	return cs.chatStore.ClearEntriesByTopic(name)

}
