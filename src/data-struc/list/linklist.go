package list

import (
	"errors"
)

type LinkedList struct {
	Size int
	Head *node
}

type node struct {
	Value interface{}
	Next  *node
}

func NewLinkedList() LinkedList {
	return LinkedList{Size: 0, Head: nil}
}

func (l *LinkedList) AddFront(value interface{}) {
	new := &node{Value: value, Next: nil}

	if n := l.Head; n != nil {
		l.Head = new
		l.Size++
		new.Next = n
		return
	}

	l.Head = new
	l.Size = 1
}

func (l *LinkedList) AddBack(value interface{}) {
	new := &node{Value: value, Next: nil}

	if l.Head == nil {
		l.Head = new
		l.Size++
		return
	}

	n := l.Head

	for n.Next != nil {
		n = n.Next
	}

	n.Next = new
	l.Size++
}

// Inserts element at index or back of the list
func (l *LinkedList) AddAtIndex(idx int, value interface{}) {
	new := &node{Value: value, Next: nil}

	if l.Head == nil {
		l.Head = new
		l.Size++
		return
	}

	var prev, cur *node
	prev, cur = nil, l.Head

	for ; idx != 0 && cur.Next != nil; idx-- {
		prev = cur
		cur = cur.Next
	}

	prev.Next = new
	new.Next = cur
	l.Size++
}

func (l *LinkedList) PopFront() (interface{}, error) {
	if l.Head == nil {
		return 0, errors.New("Cannot pop from empty list")
	}

	value := l.Head.Value
	l.Head = l.Head.Next
	l.Size--
	return value, nil
}

func (l *LinkedList) PopIdx(index int) (interface{}, error) {
	var prev, cur *node
	prev, cur = nil, l.Head

	for ; index != 0 && cur.Next != nil; index-- {
		prev = cur
		cur = cur.Next
	}

	if index != 0 {
		return 0, errors.New("Index out of range")
	}

	value := cur.Value
	prev.Next = cur.Next
	l.Size--
	return value, nil
}

func (l *LinkedList) PopBack() (interface{}, error) {
	if l.Head == nil {
		return 0, errors.New("Cannot pop from empty list")
	}

	n := l.Head

	for ; n.Next.Next != nil; n = n.Next {
	}

	r := n.Next
	n.Next = nil
	l.Size--
	return r.Value, nil
}
