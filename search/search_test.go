package search

import (
	"github.com/theosev/find-swap-go/store/memory_queue"
	"github.com/theosev/find-swap-go/user"
	"testing"
	"time"
)

func TestSearchCreation(t *testing.T) {
	mq := memory_queue.New()
	search := New(mq)

	if search == nil {
		t.Fatalf("unexpected error from New: search is empty")
	}

	if search.Queue.GetLength() > 0 {
		t.Fatalf("unexpected error from New: Queue is not empty")
	}

	if search.Timeout != 20 * time.Second {
		t.Fatalf("unexpected error from New: Timeout is not 20 seconds")
	}
}

func TestSearchNotFound(t *testing.T) {
	mq := memory_queue.New()
	search := New(mq)

	TestUser := user.New()

	notFound := search.match(TestUser)

	if notFound == true {
		t.Fatalf("unexpected error: notFound is true")
	}
}