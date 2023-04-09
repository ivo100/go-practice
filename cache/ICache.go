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

type ICache[K comparable, V any] interface {
	Get(key K) (value V, found bool)
	Set(key K, value V, duration time.Duration)
	Range(f func(key K, value V) bool)
	Delete(key K)
	Close()
}

*/
