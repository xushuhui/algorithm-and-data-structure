package datastructure

type ArrayQueue struct {
	array *Array
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		array: NewArray(10),
	}
}
