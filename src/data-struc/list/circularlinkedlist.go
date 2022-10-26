package list

import (
	"errors"
)

type CircularLinkedList struct {
	Size int
	Head *dnode
	Tail *dnode
}
type dnode struct {
	Value int
	Next  *dnode
	Prev  *dnode
}

func NewCircularLinkedList() CircularLinkedList {
	return CircularLinkedList{Size: 0, Head: nil, Tail: nil}
}

func (l *CircularLinkedList) AddFront(value int) {
	new := &dnode{Value: value, Next: nil, Prev: nil}

	if l.Head != nil {
		new.Next = l.Head
		new.Prev = l.Tail
		l.Head.Prev = new
		l.Tail.Next = new
		l.Head = new

		l.Size++
		return
	}

	l.Head = new
	l.Tail = new
	l.Size = 1
}

func (l *CircularLinkedList) AddBack(value int) {
	new := &dnode{Value: value, Next: nil}

	if l.Tail != nil {
		new.Next = l.Head
		new.Prev = l.Tail
		l.Head.Prev = new
		l.Tail.Next = new
		l.Tail = new

		l.Size++
		return
	}

	l.Head = new
	l.Tail = new
	l.Size = 1
}

// Inserts element at index or back of the list
func (l *CircularLinkedList) AddAtIndex(idx, value int) {
	new := &dnode{Value: value, Next: nil, Prev: nil}

	if l.Head == nil {
		l.Head = new
		l.Tail = new
		l.Size++
		return
	}

	if idx >= l.Size {
		l.AddBack(value)
		return
	}

	cur := l.Head

	for ; idx != 0; idx-- {
		cur = cur.Next
	}

	new.Prev = cur.Prev
	new.Next = cur
	cur.Prev.Next = new
	cur.Prev = new

	l.Size++
}

func (l *CircularLinkedList) PopFront() (int, error) {
	if l.Head == nil {
		return 0, errors.New("Cannot pop from empty list")
	}

	value := l.Head.Value
	l.remove(l.Head)
	return value, nil
}

func (l *CircularLinkedList) PopIdx(index int) (int, error) {
	cur := l.Head

	if l.Size-1 < index {
		return 0, errors.New("Index out of range")
	}

	for ; index != 0 && cur.Next != l.Tail; index-- {
		cur = cur.Next
	}

	if index == 1 {
		cur = l.Tail
	} else if index != 0 {
		return 0, errors.New("Index out of range")
	}

	value := cur.Value
	l.remove(cur)
	return value, nil
}

func (l *CircularLinkedList) PopBack() (int, error) {
	if l.Tail == nil {
		return 0, errors.New("Cannot pop from empty list")
	}

	value := l.Tail.Value
	l.remove(l.Tail)
	return value, nil
}

func (l *CircularLinkedList) remove(n *dnode) {
	prev, next := n.Prev, n.Next

	if prev != nil {
		prev.Next = next
	} else {
		l.Head = nil
	}

	if next != nil {
		next.Prev = prev
	} else {
		l.Tail = nil
	}

	if l.Head == n {
		l.Head = next
	}

	if l.Tail == n {
		l.Tail = prev
	}

	if prev == next && prev != nil {
		prev.Next = nil
		prev.Prev = nil
		next.Next = nil
		next.Prev = nil
	}

	l.Size--
}
