package cache

import "time"

type ICache interface {
	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{}, duration time.Duration)
	Range(f func(key, value interface{}) bool)
	Delete(key interface{})
	Close()
}

/*
todo:
type ICache[K comparable, V any] interface {
	Set(key K, value V, duration time.Duration)
	Get(key K) (value V, found bool)
	Delete(key K)
	Close()
	Range(visitor func(key K, value V) bool)
}
*/
