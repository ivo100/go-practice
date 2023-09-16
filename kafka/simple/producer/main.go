package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"simple/common"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	conf := common.ReadConfig(common.ConfigFile)
	p, err := kafka.NewProducer(&conf)
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	users := []string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := []string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}
	topic := common.Topic
	for n := 0; n < 10; n++ {
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(data),
		}, nil)
		if err != nil {
			log.Printf("Produce error %v", err.Error())
		}
	}

	// Wait for all messages to be delivered
	p.Flush(5 * 1000)
	p.Close()
}
