package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	a := [][]int{
		{5, 1},
		{2, 1},
		{3, 0},
		{4, 0},
		{1, 1},
		{6, 1},
	}
	fmt.Printf("%v\n", a)
	h := &IntHeap{}
	heap.Init(h)
	s := 0
	for _, p := range a {
		fmt.Printf("%v\n", p)
		if p[1] == 1 {
			heap.Push(h, p[0])
		} else {
			s += p[0]
		}
	}
	k := 3
	for k > 0 {
		n := heap.Pop(h).(int)
		fmt.Printf("k %v, %v\n", k, n)
		k--
		s += n
	}
	for h.Len() > 0 {
		n := heap.Pop(h).(int)
		s -= n
	}
	fmt.Printf("s %v\n", s)
}
