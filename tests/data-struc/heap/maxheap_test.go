package heap

import (
	"exercise-go/src/data-struc/heap"
	"exercise-go/tests"
	"testing"
)


func TestMaxHeapInsert(t *testing.T) {
	maxHeap := heap.NewMaxHeap()

	value := []int{0,1,2,3,4,5,6,7,8,9}

	for _, i := range value {
		maxHeap.Insert(i)
	}

	tests.AssertEqual(t, maxHeap.Size(), 10)
	for i := 9; i >= 0; i-- {
		tests.AssertEqual(t, maxHeap.Pop(), value[i])
	}

	tests.AssertEqual(t, maxHeap.Size(), 0)
}