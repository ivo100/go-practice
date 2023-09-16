package main

import (
	"fmt"
	"log"
	"simple/common"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

	conf := common.ReadConfig(common.ConfigFile)
	//conf["acks"] = "all"
	conf["transactional.id"] = "my-transactional-id"

	// Create a Kafka producer instance
	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		fmt.Printf("Error creating producer: %v\n", err)
		return
	}

	fmt.Println("Starting transactional producer")

	topic := common.Topic

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range producer.Events() {
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

	users := []string{"aaaa", "bbbb", "cccc"}
	items := []string{"alpha", "beta", "gamma"}

	// Initialize the producer as a transactional producer
	if err = producer.InitTransactions(nil); err != nil {
		fmt.Printf("Error initializing transactions: %v\n", err)
		return
	}

	for n := 0; n < 3; n++ {
		//key := users[rand.Intn(len(users))]
		//data := items[rand.Intn(len(items))]
		key := users[n]
		data := items[n]

		// Start a Kafka transaction
		err = producer.BeginTransaction()
		if err != nil {
			fmt.Printf("Error beginning transaction: %v\n", err)
			return
		}

		// Produce a message within the transaction
		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(data),
		}

		if err := producer.Produce(message, nil); err != nil {
			log.Printf("Produce error %v", err.Error())
			return
		}

		// Commit the Kafka transaction
		if err = producer.CommitTransaction(nil); err != nil {
			fmt.Printf("Error committing transaction: %v\n", err)
			return
		}

	}

	// Wait for all messages to be delivered
	producer.Flush(5 * 1000)
	producer.Close()
	fmt.Println("Closing transactional producer")

}
