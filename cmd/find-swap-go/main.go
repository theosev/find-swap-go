package main

import (
	"github.com/theosev/find-swap-go/store"
	"fmt"
)

func main() {
	list := store.New()

	fmt.Println(list.Dequeue())
}
