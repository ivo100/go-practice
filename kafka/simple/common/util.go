package common

import (
	"bufio"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"os"
	"strings"
)

// go get github.com/confluentinc/confluent-kafka-go/kafka
// github.com/confluentinc/confluent-kafka-go/v2/kafka
const (
	ConfigFile = "config.ini"
	Group      = "group_2"
	Topic      = "topic_2"
)

func Version() {
	vnum, vstr := kafka.LibraryVersion()
	fmt.Printf("LibraryVersion: %s (0x%x)\n", vstr, vnum)
	fmt.Printf("LinkInfo:       %s\n", kafka.LibrdkafkaLinkInfo)
}

func ReadConfig(configFile string) kafka.ConfigMap {

	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			before, after, found := strings.Cut(line, "=")
			if found {
				parameter := strings.TrimSpace(before)
				value := strings.TrimSpace(after)
				m[parameter] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return m

}
