package stack

import array2 "data-structures/datastructure/array"

type ArrayStack struct {
	array *array2.Array
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		array: array2.NewArray(10),
	}
}

func (a *ArrayStack) Push(e int) {
	a.array.AddLast(e)
}

func (a *ArrayStack) Pop() int {
	return a.array.RemoveLast()
}

func (a *ArrayStack) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *ArrayStack) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayStack) Peek() int {
	return a.array.GetLast()
}
