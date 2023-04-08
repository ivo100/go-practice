package main

import (
	"container/ring"
	"fmt"
	// relative import path - requires go mod
	// and
	// go install foo/foo.go
	//"ring/foo"
	// flaky - don't use
)

func main() {

	//foo.New()

	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < 2*n; i++ {
		r.Value = i
		r = r.Next()
	}
	// Iterate through the ring and print its contents
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

}
