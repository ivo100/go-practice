package main

import (
	"log"
	"sort"
)

// bin search of x in sorted slice a
// if found - returns index of first item == x
// if not found - returns index of first item > x
func find(a []int, x int) int {
	n := sort.SearchInts(a, x)
	// warning n == len(a) if x > max
	log.Printf("bin search %v - result idx: %v", x, n)
	return n
}

// given slice, places next element to sorted array
// shifting array 1 pos to the left
func place(a []int, x int) {
	defer log.Printf("place %v - result: %+v", x, a)
	n := find(a, x)
	// first?
	if n == 0 {
		a[0] = x
		return
	}
	// last?
	l := len(a)
	if n >= l {
		copy(a, a[1:l])
		a[l-1] = x
		return
	}
	// shift left 1:n
	// replace n
	copy(a, a[1:n])
	a[n] = x
	return
}

func main() {
	a := []int{5, 10, 2, 6, 1, 1, 2, 5, 6, 3}

	log.Printf("       a: %+v", a)
	sort.Ints(a)
	log.Printf("sorted a: %+v", a)
	x := 1
	//place(a, x)
	//x = 4
	//place(a, x)
	//x = 0
	//place(a, x)
	x = 100
	place(a, x)
}
