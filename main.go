package main

import (
	"github.com/ivo100/go-practice/cache"
	"github.com/ivo100/go-practice/userservice/userservice"
	"time"
)

type Svc struct {
	userservice.UserService
}

func main() {
	cycle := 100 * time.Millisecond
	c := cache.New(cycle)
	defer c.Close()

	c.Set("sticky", "forever", 0)
	c.Set("hello", "Hello", cycle/2)
	hello, found := c.Get("hello")

	var svc Svc
	_ = svc

}
