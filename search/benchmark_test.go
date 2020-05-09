package search

import (
	"fmt"
	"github.com/theosev/find-swap-go/store/mqueue"
	"github.com/theosev/find-swap-go/user"
	"math/rand"
	"testing"
	"time"
)

// BenchmarkFinder_Start needs improvement
func BenchmarkFinder_Start(b *testing.B) {
	mq := mqueue.New()
	search := New(mq)
	search.Timeout = time.Second
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < b.N; i++ {
		TestUser := user.New(rng.Int())
		testChannel := make(chan bool)
		defer close(testChannel)
		go search.Start(testChannel, TestUser)
		fmt.Print(<-testChannel)
	}
}

