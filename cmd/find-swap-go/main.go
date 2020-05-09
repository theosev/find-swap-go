package main

import (
	"github.com/theosev/find-swap-go/search"
	"github.com/theosev/find-swap-go/store/mqueue"
	"github.com/theosev/find-swap-go/user"
	"log"
	"math/rand"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		mq := mqueue.New()
		s := search.New(mq)
		go handleConn(conn, s)
	}

}

func handleConn(c net.Conn, s *search.Finder) {
	defer c.Close()

	u := user.New(rand.Int())
	uChannel := make(chan bool)
	defer close(uChannel)

	s.Start(uChannel, u)
	// do something
	select {
	// case received bytes from the other go routine
	// case received bytes from the client
	// timeout
	}
}
