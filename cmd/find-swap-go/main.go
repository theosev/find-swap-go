package main

import (
	"fmt"
	mq "github.com/theosev/find-swap-go/store/memory_queue"
)

func main() {
	list := mq.New()

	list.Enqueue(1)
	list.Enqueue(2)
	list.Enqueue(2)

	fmt.Println(list)
}
