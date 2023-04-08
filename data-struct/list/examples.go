package main

import (
	"container/list"
	"fmt"
)

// https://golang.org/pkg/container/list/#pkg-examples
/*
type Element
    func (e *Element) Next() *Element
    func (e *Element) Prev() *Element
type List
    func New() *List
    func (l *List) Back() *Element
    func (l *List) Front() *Element
    func (l *List) Init() *List
    func (l *List) InsertAfter(v interface{}, mark *Element) *Element
    func (l *List) InsertBefore(v interface{}, mark *Element) *Element
    func (l *List) Len() int
    func (l *List) MoveAfter(e, mark *Element)
    func (l *List) MoveBefore(e, mark *Element)
    func (l *List) MoveToBack(e *Element)
    func (l *List) MoveToFront(e *Element)
    func (l *List) PushBack(v interface{}) *Element
    func (l *List) PushBackList(other *List)
    func (l *List) PushFront(v interface{}) *Element
    func (l *List) PushFrontList(other *List)
    func (l *List) Remove(e *Element) interface{}
*/
func main() {

	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
