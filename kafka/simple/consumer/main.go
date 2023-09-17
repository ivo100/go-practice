package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"os"
	"os/signal"
	"simple/common"
	"syscall"
	"time"
)

func main() {

	conf := common.ReadConfig(common.ConfigFile)
	conf["group.id"] = common.Group
	//conf["auto.offset.reset"] = "earliest"

	c, err := kafka.NewConsumer(&conf)
	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}
	defer c.Close()

	topic := common.Topic
	err = c.SubscribeTopics([]string{topic}, nil)
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
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				fmt.Printf("Error %v - ignore\n", err.Error())
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		}
	}
}

/*

https://docs.confluent.io/kafka-clients/go/current/overview.html

// to commit async

msg_count := 0
for run == true {
    ev := consumer.Poll(100)
    switch e := ev.(type) {
    case *kafka.Message:
        msg_count += 1
        if msg_count % MIN_COMMIT_COUNT == 0 {
            go func() {
                offsets, err := consumer.Commit()
            }()
        }
        fmt.Printf("%% Message on %s:\n%s\n",
            e.TopicPartition, string(e.Value))

    case kafka.PartitionEOF:
        fmt.Printf("%% Reached %v\n", e)
    case kafka.Error:
        fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
        run = false
    default:
        fmt.Printf("Ignored %v\n", e)
    }
}

// to see rebalance events

consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
     "bootstrap.servers":    "host1:9092,host2:9092",
     "group.id":             "foo",
     "go.application.rebalance.enable": true})

msg_count := 0
for run == true {
    ev := consumer.Poll(100)
    switch e := ev.(type) {
    case kafka.AssignedPartitions:
        fmt.Fprintf(os.Stderr, "%% %v\n", e)
        c.Assign(e.Partitions)
    case kafka.RevokedPartitions:
        fmt.Fprintf(os.Stderr, "%% %v\n", e)
        c.Unassign()
    case *kafka.Message:
        msg_count += 1
        if msg_count % MIN_COMMIT_COUNT == 0 {
            consumer.Commit()
        }

        fmt.Printf("%% Message on %s:\n%s\n",
            e.TopicPartition, string(e.Value))

    case kafka.PartitionEOF:
        fmt.Printf("%% Reached %v\n", e)
    case kafka.Error:
        fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
        run = false
    default:
        fmt.Printf("Ignored %v\n", e)
    }
}

*/
