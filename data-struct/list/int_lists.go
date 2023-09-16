package main

type IntNode struct {
	Value int
	Next  *IntNode
}

type IntList interface {
	Append(val int)
	MergeSorted(a, b *IntNode) *IntNode
	GetHead() *IntNode
}

type intList struct {
	Head, Tail *IntNode
}

func NewIntList() IntList {
	return &intList{}
}

func (lst *intList) Append(val int) {
	if lst.Head == nil {
		lst.Head = &IntNode{Value: val}
		lst.Tail = lst.Head
		return
	}
	lst.Tail.Next = &IntNode{Value: val}
	lst.Tail = lst.Tail.Next
}

func (lst *intList) GetHead() *IntNode {
	return lst.Head
}

func (lst *intList) MergeSorted(a, b *IntNode) *IntNode {
	res := NewIntList()
	p, q := a, b
	for p != nil && q != nil {
		if p.Value < q.Value {
			res.Append(p.Value)
			p = p.Next
			continue
		}
		res.Append(q.Value)
		q = q.Next
	}
	for p != nil {
		res.Append(p.Value)
		p = p.Next
	}
	for q != nil {
		res.Append(q.Value)
		q = q.Next
	}
	return res.GetHead()
}

func AddIntLists(a, b *IntNode) *IntNode {
	for p, q := a, b; p != nil && q != nil; p, q = p.Next, q.Next {
		p.Value += q.Value
		if p.Next == nil || q.Next == nil {
			return a
		}
	}
	return a
}

func MergeSortedIntLists(a, b *IntNode) *IntNode {

	return a
}
