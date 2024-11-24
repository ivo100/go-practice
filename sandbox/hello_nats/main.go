package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	//simple()
	streams()
}

func simple() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Simple Publisher
	err = nc.Publish("foo", []byte("Hello 111"))
	if err != nil {
		panic(err)
	}

	// Responding to a request message
	nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})

	// Simple Sync Subscriber
	sub, err := nc.SubscribeSync("foo")
	timeout := time.Second
	m, err := sub.NextMsg(timeout)
	if err == nil {
		fmt.Printf("Received: %s\n", string(m.Data))
	}

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err = nc.ChanSubscribe("foo", ch)
	if err != nil {
		panic(err)
	}

	err = nc.Publish("foo", []byte("Hello 2222"))
	if err != nil {
		panic(err)
	}

	msg := <-ch
	fmt.Printf("Received: %s\n", string(msg.Data))
	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()

	// Replies
	// subscribe must be before publish
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Requests
	msg, err = nc.Request("help", []byte("help me"), 10*time.Millisecond)
	if msg != nil {
		fmt.Printf("Received: %s\n", string(msg.Data))
	}

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	//time.Sleep(2 * time.Second)
	// Close connection
	nc.Close()
}
