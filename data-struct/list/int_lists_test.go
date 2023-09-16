package main

import (
	"log"
	"testing"
)

func makeList(ints ...int) *IntNode {
	var head *IntNode
	var prev *IntNode
	for _, i := range ints {
		node := &IntNode{Value: i}
		if head == nil {
			head = node
		}
		if prev != nil {
			prev.Next = node
		}
		prev = node
	}
	return head
}

func printList(node *IntNode) {
	for node != nil {
		log.Println(node.Value)
		node = node.Next
	}
}

func TestAddIntLists(t *testing.T) {
	l1 := makeList(1, 1, 1)
	printList(l1)
	l2 := makeList(2, 2, 2)
	s := AddIntLists(l1, l2)
	printList(s)
}
