package main

import (
	"fmt"
	mq "github.com/theosev/find-swap-go/store/memory_queue"
)

func main() {
	list := mq.New()

	fmt.Println(list.Dequeue())
}
