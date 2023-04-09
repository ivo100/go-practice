package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Cache struct {
	data     map[string]interface{}
	ttl      time.Duration
	evictCh  chan string
	stopCh   chan struct{}
	evictMtx sync.Mutex
}

func NewCache(ttl time.Duration) *Cache {
	c := &Cache{
		data:    make(map[string]interface{}),
		ttl:     ttl,
		evictCh: make(chan string),
		stopCh:  make(chan struct{}),
	}
	go c.evictExpiredKeys()
	return c
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
	go c.scheduleEviction(key)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

func (c *Cache) evictExpiredKeys() {
	for {
		select {
		case key := <-c.evictCh:
			c.evictMtx.Lock()
			//log.Printf("Deleting key %v", key)
			delete(c.data, key)
			c.evictMtx.Unlock()
		case <-c.stopCh:
			return
		}
	}
}

func (c *Cache) scheduleEviction(key string) {
	//log.Printf("scheduleEviction key %v", key)
	time.Sleep(c.ttl)
	c.evictCh <- key
}

func (c *Cache) Stop() {
	close(c.stopCh)
}

func main() {
	debug := func() { fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine()) }

	cache := NewCache(5 * time.Second)
	defer cache.Stop()

	debug()
	cache.Set("foo", "bar")
	debug()
	time.Sleep(1 * time.Second)
	cache.Set("baz", 123)
	debug()
	time.Sleep(1 * time.Second)
	cache.Set("qux", []int{1, 2, 3})
	debug()
	time.Sleep(1 * time.Second)
	fmt.Println(cache.Get("foo")) // Output: bar
	time.Sleep(6 * time.Second)
	debug()
	fmt.Println(cache.Get("foo")) // Output: <nil> false
	time.Sleep(1 * time.Second)
	fmt.Println(cache.Get("baz")) // Output: <nil> false
	debug()

}
