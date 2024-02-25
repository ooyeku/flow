package main

import (
	"flag"
	"fmt"
	"github.com/ooyeku/flow/pkg/chat"
	"log"
)

/* Usage: go run getchathistory.go -chat cv1
 */

func main() {
	// Define a string flag for the chat version
	chatVersion := flag.String("chat", "cv1", "chat version (cv1, sv1, pv1)")

	// Parse the flags
	flag.Parse()

	// Determine the database path based on the chat version
	var dbPath string
	switch *chatVersion {
	case "cv1":
		dbPath = "internal/inmemory/cv1.db"
	case "sv1":
		dbPath = "internal/inmemory/sv1.db"
	case "pv1":
		dbPath = "internal/inmemory/pv1.db"
	default:
		log.Fatalf("invalid chat version: %s", *chatVersion)
	}

	chatStore, err := chat.NewChatStore(dbPath)
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
