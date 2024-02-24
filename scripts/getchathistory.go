package main

import (
	"fmt"
	"github.com/ooyeku/flow/pkg/chat"
	"log"
)

func main() {
	chatStore, err := chat.NewChatStore("chat.db")
	if err != nil {
		log.Fatalf("error creating chat store: %s", err)
	}
	defer func(chatStore *chat.ChatStore) {
		err := chatStore.Close()
		if err != nil {
			log.Fatalf("error closing chat store: %s", err)
		}
	}(chatStore)

	entries, err := chatStore.RetrieveEntries()
	if err != nil {
		log.Fatalf("error retrieving chat entries: %s", err)
	}

	for _, entry := range entries {
		fmt.Printf("%s: %s\n", entry.Timestamp.Format("2006-01-02 15:04:05"), entry.UserInput)
		fmt.Printf("AI: %s\n", entry.AIResponse)
	}
}
