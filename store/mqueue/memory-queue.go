package mqueue

// MemoryQueue type
type MemoryQueue struct {
	list []interface{}
}

// New creates a new queue
func New() *MemoryQueue {
	return &MemoryQueue{}
}

// Enqueue adds an element to the list
func (q *MemoryQueue) Enqueue(element interface{}) {
	q.list = append(q.list, element)
}

// Dequeue removes an element from the list
func (q *MemoryQueue) Dequeue() interface{} {
	if len(q.list) == 0 {
		return 0
	}

	element := q.list[0]
	q.list = q.list[1:]
	return element
}

// GetLength returns the length of the queue
func (q *MemoryQueue) GetLength() int {
	return len(q.list)
}
