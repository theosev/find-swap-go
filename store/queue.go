package store

type Queue interface {
	Enqueue(element interface{})
	Dequeue() interface{}
}