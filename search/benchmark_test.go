package search

import (
	"github.com/theosev/find-swap-go/store/mqueue"
	"github.com/theosev/find-swap-go/user"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkFinder_Start(b *testing.B) {
	mq := mqueue.New()
	search := New(mq)
	search.Timeout = time.Second

	for i := 0; i < b.N; i++ {
		TestUser := user.New(rand.Intn(1000000))
		testChannel := make(chan bool)
		defer close(testChannel)
		go search.Start(testChannel, TestUser)
		<-testChannel
	}
}

