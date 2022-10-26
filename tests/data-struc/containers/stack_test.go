package containers_test

import (
	"exercise-go/src/data-struc/containers"
	"exercise-go/tests"
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := containers.NewStack()

	tests.AssertEqual(t, s.Size(), 0)

	s.Push(1)
	tests.AssertEqual(t, s.Size(), 1)
	value, err := s.Pop()
	tests.AssertEqual(t, s.Size(), 0)
	tests.AssertEqual(t, value, 1)
	tests.AssertEqual(t, err, nil)

	value, err = s.Pop()
	tests.AssertEqual(t, fmt.Sprint(err), "Cannot pop from empty Stack")
}
