package chat

import (
	_ "github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	_ "time"
)

func setupChatStore(t *testing.T) *ChatStore {
	cs, err := NewChatStore("test.db")
	if err != nil {
		t.Fatalf("failed to create chat store: %v", err)
	}

	return cs
}

func teardownChatStore(t *testing.T, cs *ChatStore) {
	err := cs.db.Select(q.True()).Delete(new(ChatEntry))
	if err != nil {
		t.Fatalf("failed to delete entries: %v", err)
	}

	err = cs.Close()
	if err != nil {
		t.Fatalf("failed to close chat store: %v", err)
	}

	err = os.Remove("test.db")

}

func TestChatStore_SaveEntry(t *testing.T) {
	cs := setupChatStore(t)
	defer teardownChatStore(t, cs)

	err := cs.SaveEntry("hello", "hi")
	if err != nil {
		t.Fatalf("failed to save entry: %v", err)
	}

	entries, err := cs.RetrieveEntries()
	if err != nil {
		t.Fatalf("failed to retrieve entries: %v", err)
	}

	assert.Equal(t, 1, len(entries))
	assert.Equal(t, "hello", entries[0].UserInput)
	assert.Equal(t, "hi", entries[0].AIResponse)
}

func TestChatStore_RetrieveEntries(t *testing.T) {
	cs := setupChatStore(t)
	defer teardownChatStore(t, cs)

	err := cs.SaveEntry("hello", "hi")
	if err != nil {
		t.Fatalf("failed to save entry: %v", err)
	}

	entries, err := cs.RetrieveEntries()
	if err != nil {
		t.Fatalf("failed to retrieve entries: %v", err)
	}

	assert.Equal(t, 1, len(entries))
	assert.Equal(t, "hello", entries[0].UserInput)
	assert.Equal(t, "hi", entries[0].AIResponse)
}
