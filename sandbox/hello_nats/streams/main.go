package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	// In the `jetstream` package, almost all API calls rely on `context.Context` for timeout/cancellation handling
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	nc, _ := nats.Connect(nats.DefaultURL)

	// Create a JetStream management interface
	js, _ := jetstream.New(nc)

	// Create a stream
	s, _ := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})

	// Publish some messages
	for i := 0; i < 100; i++ {
		js.Publish(ctx, "ORDERS.new", []byte("hello message "+strconv.Itoa(i)))
		fmt.Printf("Published hello message %d\n", i)
	}

	// Create durable consumer
	c, _ := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "CONS",
		AckPolicy: jetstream.AckExplicitPolicy,
	})

	// Get 10 messages from the consumer
	messageCounter := 0
	msgs, err := c.Fetch(10)
	if err != nil {
		// handle error
	}

	for msg := range msgs.Messages() {
		msg.Ack()
		fmt.Printf("Received a JetStream message via fetch: %s\n", string(msg.Data()))
		messageCounter++
	}

	fmt.Printf("received %d messages\n", messageCounter)

	if msgs.Error() != nil {
		fmt.Println("Error during Fetch(): ", msgs.Error())
	}

	// Receive messages continuously in a callback
	cons, _ := c.Consume(func(msg jetstream.Msg) {
		msg.Ack()
		fmt.Printf("Received a JetStream message via callback: %s\n", string(msg.Data()))
		messageCounter++
	})
	defer cons.Stop()

	// Iterate over messages continuously
	it, _ := c.Messages()
	for i := 0; i < 10; i++ {
		msg, _ := it.Next()
		msg.Ack()
		fmt.Printf("Received a JetStream message via iterator: %s\n", string(msg.Data()))
		messageCounter++
	}
	it.Stop()

	// block until all 100 published messages have been processed
	for messageCounter < 100 {
		time.Sleep(10 * time.Millisecond)
	}
}
