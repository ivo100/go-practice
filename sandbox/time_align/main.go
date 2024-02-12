package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Hello")

	quit := make(chan bool)
	data := make(chan string)
	minute := time.Second * 60
	// Work at aligned time intervals - e.g. top of every minute
	// until quit message is received
	go func() {
		for {
			now := time.Now()
			delay := now.Truncate(minute).Add(minute).Sub(now)

			select {
			case <-quit:
				log.Println("Quit")
				close(data)
				return
			case <-time.After(delay):
				data <- fetch()
			}
		}
	}()

	go func() {
		time.Sleep(120 * time.Second)
		log.Println("Sending quit message")
		quit <- true
	}()

loop:
	for {
		select {
		case info, ok := <-data:
			if !ok {
				break loop
			}
			log.Println("Working...", info)
		}
	}

	log.Println("Goodbye")
}

func fetch() string {
	//log.Println("Fetching")
	return "some work"
}
