package containers

import (
	"exercise-go/src/data-struc/list"
	"errors"
)

type Queue struct {
	data list.CircularLinkedList
}

func NewQueue() Queue {
	return Queue{data: list.NewCircularLinkedList()}
}

func (q *Queue) Push(value interface{}) {
	q.data.AddBack(value)
}

func (q *Queue) Pop() (interface{}, error) {
	v, e := q.data.PopFront()

	if e != nil {
		e = errors.New("Cannot pop from empty Queue")
	}

	return v, e
}

func (q *Queue) Size() int {
	return q.data.Size
}
