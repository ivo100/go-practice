package main

import (
	"log"
	"strings"
)

func main() {
	s := "Hello world"

	idx := strings.Index(s, "world")

	log.Printf("idx %v", idx)
}
