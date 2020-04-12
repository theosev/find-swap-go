package main

import (
	"find-swap-go/store"
	"fmt"
)

func main() {
	list := store.New()

	fmt.Println(list.Dequeue())
}
