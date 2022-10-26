package list_test

import (
	"exercise-go/src/data-struc/list"
	"exercise-go/tests"
	"fmt"
	"testing"
)

func TestCircularLinkedList_AddBack(t *testing.T) {
	l := list.NewCircularLinkedList()

	tests.AssertEqual(t, l.Size, 0)
	// utils.go:14: Received <nil> (type *list.dnode), expected <nil> (type <nil>)
	// tests.AssertEqual(t, l.Head, nil)
	// utils.go:14: Received <nil> (type *list.dnode), expected <nil> (type <nil>)
	// tests.AssertEqual(t, l.Tail, nil)

	// [0]
	l.AddBack(0)
	tests.AssertEqual(t, l.Size, 1)
	tests.AssertEqual(t, l.Head.Value, 0)
	tests.AssertEqual(t, l.Tail.Value, 0)

	// [0 1]
	l.AddBack(1)
	tests.AssertEqual(t, l.Size, 2)
	tests.AssertEqual(t, l.Head.Value, 0)
	tests.AssertEqual(t, l.Head.Next.Value, 1)
	tests.AssertEqual(t, l.Tail.Value, 1)
	tests.AssertEqual(t, l.Tail.Prev.Value, 0)

	// [0 1 2]
	l.AddBack(2)
	tests.AssertEqual(t, l.Size, 3)
	tests.AssertEqual(t, l.Head.Value, 0)
	tests.AssertEqual(t, l.Tail.Value, 2)
	tests.AssertEqual(t, l.Head.Next.Value, 1)
	tests.AssertEqual(t, l.Head.Prev.Value, 2)
	tests.AssertEqual(t, l.Tail.Next.Value, 0)
	tests.AssertEqual(t, l.Tail.Prev.Value, 1)
}

func TestCircularLinkedList_AddFront(t *testing.T) {
	l := list.NewCircularLinkedList()

	tests.AssertEqual(t, l.Size, 0)
	// utils.go:14: Received <nil> (type *list.dnode), expected <nil> (type <nil>)
	// tests.AssertEqual(t, l.Head, nil)
	// utils.go:14: Received <nil> (type *list.dnode), expected <nil> (type <nil>)
	// tests.AssertEqual(t, l.Tail, nil)

	// [0]
	l.AddFront(0)
	tests.AssertEqual(t, l.Size, 1)
	tests.AssertEqual(t, l.Head.Value, 0)
	tests.AssertEqual(t, l.Tail.Value, 0)

	// [1 0]
	l.AddFront(1)
	tests.AssertEqual(t, l.Size, 2)
	tests.AssertEqual(t, l.Head.Value, 1)
	tests.AssertEqual(t, l.Head.Prev.Value, 0)
	tests.AssertEqual(t, l.Head.Next.Value, 0)
	tests.AssertEqual(t, l.Tail.Value, 0)
	tests.AssertEqual(t, l.Tail.Prev.Value, 1)
	tests.AssertEqual(t, l.Tail.Next.Value, 1)

	// [2 1 0]
	l.AddFront(2)
	tests.AssertEqual(t, l.Size, 3)
	tests.AssertEqual(t, l.Head.Value, 2)
	tests.AssertEqual(t, l.Tail.Value, 0)
	tests.AssertEqual(t, l.Head.Next.Value, 1)
	tests.AssertEqual(t, l.Head.Prev.Value, 0)
	tests.AssertEqual(t, l.Tail.Prev.Value, 1)
	tests.AssertEqual(t, l.Tail.Next.Value, 2)
}

func TestCircularLinkedList_AddAtIndex(t *testing.T) {
	l := list.NewCircularLinkedList()

	tests.AssertEqual(t, l.Size, 0)

	// [1]
	l.AddAtIndex(1, 1)
	tests.AssertEqual(t, l.Size, 1)
	tests.AssertEqual(t, l.Head.Value, 1)
	tests.AssertEqual(t, l.Tail.Value, 1)

	// [1 3]
	l.AddAtIndex(1, 3)
	tests.AssertEqual(t, l.Size, 2)
	tests.AssertEqual(t, l.Head.Value, 1)
	tests.AssertEqual(t, l.Head.Next.Value, 3)
	tests.AssertEqual(t, l.Head.Prev.Value, 3)
	tests.AssertEqual(t, l.Tail.Value, 3)
	tests.AssertEqual(t, l.Tail.Next.Value, 1)
	tests.AssertEqual(t, l.Tail.Prev.Value, 1)

	// [1 2 3]
	l.AddAtIndex(1, 2)
	tests.AssertEqual(t, l.Size, 3)
	tests.AssertEqual(t, l.Head.Value, 1)
	tests.AssertEqual(t, l.Head.Next.Value, 2)
	tests.AssertEqual(t, l.Head.Prev.Value, 3)
	tests.AssertEqual(t, l.Tail.Value, 3)
	tests.AssertEqual(t, l.Tail.Next.Value, 1)
	tests.AssertEqual(t, l.Tail.Prev.Value, 2)
}

func TestCircularLinkedList_PopFront(t *testing.T) {
	l := list.NewCircularLinkedList()
	l.AddBack(0)
	l.AddBack(1)
	l.AddBack(2)

	// 0 <- [1 2]
	value, err := l.PopFront()
	tests.AssertEqual(t, l.Size, 2)
	tests.AssertEqual(t, value, 0)

	// 1 <- [2]
	value, err = l.PopFront()
	tests.AssertEqual(t, l.Size, 1)
	tests.AssertEqual(t, value, 1)

	// 2 <- []
	value, err = l.PopFront()
	tests.AssertEqual(t, l.Size, 0)
	tests.AssertEqual(t, value, 2)

	// err <- []
	value, err = l.PopFront()
	tests.AssertEqual(t, fmt.Sprint(err), "Cannot pop from empty list")
}

func TestCircularLinkedList_PopBack(t *testing.T) {
	l := list.NewCircularLinkedList()
	l.AddBack(0)
	l.AddBack(1)
	l.AddBack(2)

	// 2 <- [0 1]
	value, err := l.PopBack()
	tests.AssertEqual(t, l.Size, 2)
	tests.AssertEqual(t, value, 2)

	// 1 <- [0]
	value, err = l.PopBack()
	tests.AssertEqual(t, l.Size, 1)
	tests.AssertEqual(t, value, 1)

	// 0 <- []
	value, err = l.PopBack()
	tests.AssertEqual(t, l.Size, 0)
	tests.AssertEqual(t, value, 0)

	// err <- []
	value, err = l.PopBack()
	tests.AssertEqual(t, fmt.Sprint(err), "Cannot pop from empty list")
}


func TestCircularLinkedList_PopIdx(t *testing.T) {
	l := list.NewCircularLinkedList()
	l.AddBack(0)
	l.AddBack(1)
	l.AddBack(2)
	l.AddBack(3)
	l.AddBack(4)
	l.AddBack(5)

	// 0 <- [1 2 3 4 5]
	value, err := l.PopIdx(0)
	tests.AssertEqual(t, l.Size, 5)
	tests.AssertEqual(t, value, 0)

	// 5 <- [1 2 3 4]
	value, err = l.PopIdx(l.Size - 1)
	tests.AssertEqual(t, l.Size, 4)
	tests.AssertEqual(t, value, 5)

	// 3 <- [1 2 4]
	value, err = l.PopIdx(l.Size / 2)
	tests.AssertEqual(t, l.Size, 3)
	tests.AssertEqual(t, value, 3)

	// err <- [1 2 4]
	value, err = l.PopIdx(4)
	tests.AssertEqual(t, fmt.Sprint(err), "Index out of range")


	l = list.NewCircularLinkedList()
	value, err = l.PopIdx(4)
	tests.AssertEqual(t, fmt.Sprint(err), "Index out of range")
}
