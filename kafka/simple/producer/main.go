package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"simple/common"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	common.Version() //
	conf := common.ReadConfig(common.ConfigFile)
	p, err := kafka.NewProducer(&conf)
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// log message delivery reports and
	// possibly other event types (errors, stats, etc)
	go common.LogDelivery(p)
	//deliveryChan := make(chan kafka.Event, 10000)
	users := []string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := []string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}
	topic := common.Topic
	for n := 0; n < 3; n++ {
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(data),
		}
		if err := p.Produce(msg, nil); err != nil {
			log.Printf("Produce error %v", err.Error())
		}
		time.Sleep(1 * time.Second)
	}

	// Wait for all messages to be delivered
	p.Flush(3 * 1000)
	p.Close()
}
