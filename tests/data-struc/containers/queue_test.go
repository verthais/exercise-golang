package containers_test

import (
	"exercise-go/src/data-struc/containers"
	"exercise-go/tests"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := containers.NewQueue()

	tests.AssertEqual(t, q.Size(), 0)

	q.Push(1)
	tests.AssertEqual(t, q.Size(), 1)
	value, err := q.Pop()
	tests.AssertEqual(t, q.Size(), 0)
	tests.AssertEqual(t, value, 1)
	tests.AssertEqual(t, err, nil)

	value, err = q.Pop()
	tests.AssertEqual(t, fmt.Sprint(err), "Cannot pop from empty Queue")
}
