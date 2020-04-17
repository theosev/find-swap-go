package search

import (
	"github.com/theosev/find-swap-go/store/memory_queue"
	"time"
)

const Timeout  = 20 * time.Second
type Search struct {
	Queue *memory_queue.MemoryQueue
	Timeout time.Duration
}

func New(mq *memory_queue.MemoryQueue) *Search {
	return &Search{
		Queue:mq,
		Timeout: Timeout,
	}
}

func (s *Search) match(object interface{})  bool {
	return true
}

func (s *Search) logic(element interface{}) {
	// Lock
	if s.Queue.GetLength() > 0 {
		s.Queue.Dequeue()
		// send signal
	}

	s.Queue.Enqueue(element)
	// Unlock
}