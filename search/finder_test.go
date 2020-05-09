package search

import (
	"github.com/theosev/find-swap-go/store/mqueue"
	"github.com/theosev/find-swap-go/user"
	"testing"
	"time"
)

func TestSearchCreation(t *testing.T) {
	mq := mqueue.New()
	search := New(mq)

	if search == nil {
		t.Fatalf("unexpected error from New: search is empty")
	}

	if search.Queue.GetLength() > 0 {
		t.Fatalf("unexpected error from New: Queue is not empty")
	}

	if search.Timeout != timeout {
		t.Fatalf("unexpected error from New: Timeout is not 20 seconds")
	}
}

func TestFinder_NotFound(t *testing.T) {
	mq := mqueue.New()
	search := New(mq)
	search.Timeout = time.Second

	TestUser := user.New(1)

	testChannel := make(chan bool)
	go search.Start(testChannel, TestUser)

	result := <-testChannel
	//defer close(testChannel)
	if result == true {
		t.Fatalf("unexpected error: notFound is true")
	}
}

func TestFinder_Found(t *testing.T) {
	mq := mqueue.New()
	search := New(mq)
	search.Timeout = time.Second

	// Create first user
	TestUser1 := user.New(1)
	testChannel1 := make(chan bool)
	defer close(testChannel1)
	go search.Start(testChannel1, TestUser1)

	// Create second user
	TestUser2 := user.New(2)
	testChannel2 := make(chan bool)
	defer close(testChannel2)
	go search.Start(testChannel2, TestUser2)

	user2 := <-testChannel2
	user1 := <-testChannel1

	if user2 == false {
		t.Fatalf("unexpected error: notFound is true")
	}

	if user1 == false {
		t.Fatalf("unexpected error: notFound is true")
	}

}

func TestFinder_MoreUsers(t *testing.T) {
	var expectedResult []bool

	mq := mqueue.New()
	search := New(mq)
	search.Timeout = time.Second

	// Create first user
	TestUser1 := user.New(1)
	testChannel1 := make(chan bool)
	defer close(testChannel1)
	go search.Start(testChannel1, TestUser1)

	// Create second user
	TestUser2 := user.New(2)
	testChannel2 := make(chan bool)
	defer close(testChannel2)
	go search.Start(testChannel2, TestUser2)

	// Create third user
	TestUser3 := user.New(3)
	testChannel3 := make(chan bool)
	defer close(testChannel3)
	go search.Start(testChannel3, TestUser3)

	user1 := <-testChannel1
	user2 := <-testChannel2


	user3 := <-testChannel3

	expectedResult = append(expectedResult, user1)
	expectedResult = append(expectedResult, user2)
	expectedResult = append(expectedResult, user3)

	var numberFalse int
	var numberTrue int

	for _, value := range expectedResult {
		if value == false {
			numberFalse++
		}
		if value == true {
			numberTrue++
		}
	}

	if numberFalse != 1 {
		t.Fatalf("unexpected error: numberFalse should be 1: %d", numberFalse)
	}

	if numberTrue != 2 {
		t.Fatalf("unexpected error: numberTrue should be 2: %d", numberTrue)
	}
}