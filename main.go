package main

import (
	"fmt"
	"github.com/ivo100/go-practice/cache"
	userservice "github.com/ivo100/go-practice/userservice/pkg"
	"log"
	"time"
)

//

type Svc struct {
	userservice.UserService
}

type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

func (tz TimeZone) String() string {
	return fmt.Sprintf("UTC%+d", tz)
}

func main() {
	duration := 100 * time.Millisecond
	c := cache.New(duration)
	defer c.Close()

	log.Printf("EST %v, PST %v", EST, PST)

	c.Set("sticky", "forever", 0)
	c.Set("hello", "Hello", duration/2)
	hello, found := c.Get("hello")
	log.Printf("cache TTL %v", duration)
	log.Printf("found %v, value %v", found, hello)
	time.Sleep(duration)
	hello, found = c.Get("hello")
	log.Printf("must expire - found %v, value %v", found, hello)

	var svc Svc
	_ = svc
}
