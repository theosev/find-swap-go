package search

import (
	"github.com/theosev/find-swap-go/store/mqueue"
	"github.com/theosev/find-swap-go/user"
	"sync"
	"time"
)

const timeout = 20 * time.Second

// Finder struct
type Finder struct {
	Queue   *mqueue.MemoryQueue
	Timeout time.Duration
	Mutex   sync.Mutex
}

// Element struct
type Element struct {
	Object        interface{}
	ObjectChannel chan bool
}

// New creates a new Finder
func New(mq *mqueue.MemoryQueue) *Finder {
	return &Finder{
		Queue:   mq,
		Timeout: timeout,
	}
}

// Start takes an object and match it with another one in the list
func (f *Finder) Start(result chan bool, object *user.User) {
	c := make(chan bool)
	defer close(c)

	element := Element{
		Object:        object,
		ObjectChannel: c,
	}

	go match(f, element)
	select {
	case found := <-element.ObjectChannel:
		result <- found
	case <-time.After(f.Timeout):
		f.Mutex.Lock()
		f.Queue.Dequeue()
		f.Mutex.Unlock()
		result <- false
	}
}

// match is a go routine
func match(f *Finder, element Element) {
	// Lock
	if f.Queue.GetLength() > 0 {
		f.Mutex.Lock()
		queueElement := f.Queue.Dequeue()
		f.Mutex.Unlock()
		queueUser := queueElement.(Element)
		queueUser.ObjectChannel <- true
		element.ObjectChannel <- true
		return
	}

	f.Mutex.Lock()
	f.Queue.Enqueue(element)
	f.Mutex.Unlock()
}
