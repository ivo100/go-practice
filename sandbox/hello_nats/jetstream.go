package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

// https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

func streams() {

	// In the `jetstream` package, almost all API calls rely on `context.Context` for timeout/cancellation handling
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	nc, _ := nats.Connect(nats.DefaultURL)

	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}

	// Create a JetStream management interface
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}
	s, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "TEST_STREAM",
		Subjects: []string{"FOO.*"},
	})
	if err != nil {
		log.Fatal(err)
	}
	// Create durable consumer
	cons, err := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "TestConsumerParallelConsume",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal(err)
	}
	go endlessPublish(ctx, nc, js)
	for i := 0; i < 5; i++ {
		// Receive messages continuously in a callback
		cc, err := cons.Consume(func(consumeID int) jetstream.MessageHandler {
			return func(msg jetstream.Msg) {
				fmt.Printf("Received a JetStream message via callback: %s\n", string(msg.Data()))
				msg.Ack()
			}
		}(i), jetstream.ConsumeErrHandler(func(consumeCtx jetstream.ConsumeContext, err error) {
			fmt.Println(err)
		}))
		if err != nil {
			log.Fatal(err)
		}
		defer cc.Stop()
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}

func endlessPublish(ctx context.Context, nc *nats.Conn, js jetstream.JetStream) {
	var i int
	for {
		time.Sleep(500 * time.Millisecond)
		if nc.Status() != nats.CONNECTED {
			continue
		}
		if _, err := js.Publish(ctx, "FOO.TEST1", []byte(fmt.Sprintf("msg %d", i))); err != nil {
			fmt.Println("pub error: ", err)
		}
		i++
	}
}
