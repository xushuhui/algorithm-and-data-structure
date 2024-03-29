package queue

import (
	"data-structures/datastructure/array"
)

type ArrayQueue struct {
	array *array.Array
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		array: array.NewArray(10),
	}
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *ArrayQueue) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayQueue) EnQueue(e int) {
	a.array.AddLast(e)
}

func (a *ArrayQueue) DeQueue() int {
	return a.array.RemoveFirst()
}

func (a *ArrayQueue) Front() int {
	return a.array.GetFirst()
}
