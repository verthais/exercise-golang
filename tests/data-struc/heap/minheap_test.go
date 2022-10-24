package heap

import (
	"exercise-go/src/data-struc/heap"
	"exercise-go/tests"
	"testing"
)


func TestMinHeapInsert(t *testing.T) {
	maxHeap := heap.NewMinHeap()

	value := []int{9,8,7,6,5,4,3,2,1,0}

	for _, i := range value {
		maxHeap.Insert(i)
	}

	tests.AssertEqual(t, maxHeap.Size(), 10)
	for i := 9; i >= 0; i-- {
		tests.AssertEqual(t, maxHeap.Pop(), value[i])
	}

	tests.AssertEqual(t, maxHeap.Size(), 0)
}