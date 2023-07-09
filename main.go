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

//----------------------------------------------------------------------------

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
	c.Set("val", "Hello", duration/2)
	val, found := c.Get("val")
	log.Printf("cache TTL %v", duration)
	log.Printf("found %v, value %v", found, val)
	time.Sleep(duration)
	val, found = c.Get("val")
	log.Printf("must expire - found %v, value %v", found, val)
	val, found = c.Get("sticky")
	log.Printf("sticky - found %v, value %v", found, val)

	var svc Svc
	_ = svc
}
