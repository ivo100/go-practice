package trees

// Problem statement:
// You are given the roots of two binary trees and must determine if these trees are identical.
// Identical trees have the same layout and data at each node.

type Node[T comparable] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func Equals[T comparable](m, n *Node[T]) bool {
	if m == nil && n == nil {
		return true
	}
	if m == nil && n != nil {
		return false
	}
	if m != nil && n == nil {
		return false
	}
	if m.Value != n.Value {
		return false
	}
	if !Equals(m.Left, n.Left) {
		return false
	}
	if !Equals(m.Right, n.Right) {
		return false
	}
	return true
}
