package datastructure

type ArrayStack struct {
	array *Array
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		array: NewArray(10),
	}
}
func (a *ArrayStack) Push(e interface{}) {
	a.array.AddLast(e)
}
func (a *ArrayStack) Pop() interface{} {
	a.array.RemoveLast()
	return nil
}
func (a *ArrayStack) IsEmpty() bool {
	return a.array.IsEmpty()
}
func (a *ArrayStack) GetSize() int {
	return a.array.GetSize()
}
func (a *ArrayStack) Peek() interface{} {
	a.array.GetLast()
	return nil
}
