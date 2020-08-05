package datastructure

type ArrayQueue struct {
	array *Array
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		array: NewArray(10),
	}
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.array.IsEmpty()
}
func (a *ArrayQueue) GetSize() int {
	return a.array.GetSize()
}
func (a *ArrayQueue) EnQueue(e interface{}) {
	a.array.AddLast(e)
}
func (a *ArrayQueue) DeQueue() interface{} {
	return a.array.RemoveFirst()
}
func (a *ArrayQueue) Front() interface{} {
	return a.array.GetFirst()
}
