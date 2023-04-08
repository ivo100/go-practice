package main

import (
	"fmt"
	"log"
	"strings"
)

/*
https://practice.geeksforgeeks.org/problems/deletion-and-reverse-in-linked-list/1

Given a Circular Linked List of size N.
The task is to delete the given node (excluding the first and last node)
in circular linked list and then print the reverse of the circular linked list.

2 5 7 8 10
8

1 7 8 10
8

Output:
10 7 5 2
10 7 1
*/
func main() {
	a := []int{1, 7, 8, 10}
	//a := []int{1, 2, 3}
	c := NewList(a)
	log.Printf("list %v", c)
	c.Reverse()
	log.Printf(" rev %v", c)
	c.Delete(8)
	log.Printf("list %v", c)
	c.Reverse()
	log.Printf(" rev %v", c)
}

type Node struct {
	V    int
	Next *Node
}

type List struct {
	head *Node
	len  int
}

func NewList(a []int) *List {
	lst := &List{}
	if len(a) == 0 {
		return lst
	}
	lst.len = len(a)
	lst.head = &Node{V: a[0]}
	prev := lst.head
	for i := 1; i < len(a); i++ {
		x := &Node{V: a[i]}
		prev.Next = x
		prev = x
	}
	prev.Next = lst.head
	return lst
}

func (l *List) Delete(x int) bool {
	if l.head == nil {
		return false
	}
	n := l.head
	prev := n
	n = n.Next
	for {
		if n.V == x {
			prev.Next = n.Next
			n = nil
			l.len--
			return true
		}
		prev = n
		if n == l.head {
			break
		}
		n = n.Next
	}
	return false
}

func (l *List) Reverse() {
	if l.head == nil {
		return
	}

	// 1. Initialize three pointers prev as NULL,
	//    curr as head and next as NULL.
	// 2. Iterate trough the linked list. In loop, do following.
	// Before changing next of current,
	// store next node
	//next = curr->next

	// Now change next of current
	// This is where actual reversing happens
	//curr->next = prev

	// Move prev and curr one step forward
	//prev = curr
	//curr = next

	var first, prev, curr, next *Node
	first = l.head
	prev = first
	curr = first.Next
	for {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		if curr == l.head {
			l.head = prev
			first.Next = l.head
			break
		}
	}
}

func (l *List) String() string {
	if l.head == nil {
		return "<empty>"
	}
	var sb strings.Builder
	n := l.head
	for i := 0; i < l.len; i++ {
		//sb.WriteString(fmt.Sprintf("%+v ", n))
		sb.WriteString(fmt.Sprintf("%v ", n.V))
		n = n.Next
	}
	return sb.String()
}
