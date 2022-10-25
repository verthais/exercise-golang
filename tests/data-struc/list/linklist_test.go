package list

import (
	"exercise-go/src/data-struc/list"
	"exercise-go/tests"
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	// []
	l := list.NewLinkedList()
	tests.AssertEqual(t, l.Size, 0)
	// Error: utils.go:14: Received <nil> (type *list.node), expected <nil> (type <nil>)
	// tests.AssertEqual(t, l.Head, nil)

	value, err := l.PopFront()
	tests.AssertEqual(t, fmt.Sprint(err), "Cannot pop from empty list")

	value, err = l.PopBack()
	tests.AssertEqual(t, fmt.Sprint(err), "Cannot pop from empty list")

	// [1]
	l.AddBack(1)
	tests.AssertEqual(t, l.Size, 1)
	tests.AssertEqual(t, l.Head.Value, 1)

	// [1, 3]
	l.AddBack(3)
	tests.AssertEqual(t, l.Size, 2)
	tests.AssertEqual(t, l.Head.Value, 1)
	tests.AssertEqual(t, l.Head.Next.Value, 3)

	// [0, 1, 3]
	l.AddFront(0)
	tests.AssertEqual(t, l.Size, 3)
	tests.AssertEqual(t, l.Head.Value, 0)
	tests.AssertEqual(t, l.Head.Next.Value, 1)

	// [0, 1, 2, 3]
	l.AddIndex(2, 2)
	tests.AssertEqual(t, l.Size, 4)

	// [0, 1, 2, 3]
	index, node := list.Find(&l, 2)
	tests.AssertEqual(t, index, 2)
	tests.AssertEqual(t, node.Value, 2)

	// [0, 1, 2, 3]
	index, node = list.Find(&l, 3)
	tests.AssertEqual(t, index, 3)
	tests.AssertEqual(t, node.Value, 3)
	// Error: utils.go:14: Received <nil> (type *list.node), expected <nil> (type <nil>)
	// tests.AssertEqual(t, node, nil)

	// 0 <- [1, 2, 3]
	value, err = l.PopFront()
	tests.AssertEqual(t, err, nil)
	tests.AssertEqual(t, 0, value)
	tests.AssertEqual(t, l.Size, 3)

	// 2 <- [1, 3]
	value, err = l.PopIdx(1)
	tests.AssertEqual(t, err, nil)
	tests.AssertEqual(t, value, 2)
	tests.AssertEqual(t, l.Size, 2)

	value, err = l.PopIdx(9)
	tests.AssertEqual(t, fmt.Sprint(err), "Index out of range")

	// 3 <- [1]
	value, err = l.PopBack()
	tests.AssertEqual(t, err, nil)
	tests.AssertEqual(t, l.Size, 1)
	tests.AssertEqual(t, value, 3)
}
