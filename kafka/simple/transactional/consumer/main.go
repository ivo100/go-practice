package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"simple/common"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting transactional consumer")
	conf := common.ReadConfig(common.ConfigFile)
	conf["group.id"] = common.Group
	//conf["auto.offset.reset"] = "earliest"
	conf["enable.auto.commit"] = false

	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	topic := common.Topic
	err = consumer.SubscribeTopics([]string{topic}, nil)
	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
			_, err = consumer.CommitMessage(ev)
			if err != nil {
				fmt.Printf("Error committing transaction: %v\n", err)
				continue
			}

		}
	}
	fmt.Println("Closing transactional consumer")
	consumer.Close()
}
