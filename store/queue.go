package store

// Queue interface
type Queue interface {
	Enqueue(element interface{})
	Dequeue() interface{}
}
