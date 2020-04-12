package memory_queue_test

import (
	mqueue "find-swap-go/store/memory_queue"
	"testing"
)

func TestNewQueue(t *testing.T) {
	store := mqueue.New()

	if store == nil {
		t.Fatalf("unexpected error from New: Queue is empty")
	}

}

func TestQueueOperations(t *testing.T) {
	testValue := 5
	store := mqueue.New()


	store.Enqueue(testValue)

	value := store.Dequeue()

	if value != testValue {
		t.Fatalf("unexpected error test value: %d, actual value: %d", testValue, value)
	}

	emptyValue := store.Dequeue()

	if emptyValue != 0 {
		t.Fatalf("unexpected error: queue is not empty")
	}

}
