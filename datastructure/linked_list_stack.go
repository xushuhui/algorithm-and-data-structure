package datastructure

type Stack interface {
	Push(e interface{})
	Pop() interface{}
	IsEmpty() bool
	GetSize() int
	Peek() interface{}
}
type LinkedListStack struct {
	linkedlist *LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		linkedlist: NewLinkedList(),
	}
}

func (l *LinkedListStack) Push(e interface{}) {
	l.linkedlist.AddFirst(e)
}

func (l *LinkedListStack) Pop() interface{} {
	return l.linkedlist.RemoveFirst()
}

func (l *LinkedListStack) Peek() interface{} {
	return l.linkedlist.GetFirst()
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.linkedlist.IsEmpty()
}

func (l *LinkedListStack) GetSize() int {
	return l.linkedlist.GetSize()
}

func (l *LinkedListStack) Print() *LinkedList {
	return l.linkedlist
}
