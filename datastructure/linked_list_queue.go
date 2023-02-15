package datastructure

type Queue interface {
	IsEmpty() bool
	GetSize() int
	EnQueue(e interface{})
	DeQueue() interface{}
	Front() interface{}
}
type LinkedListQueue struct {
	linkedlist *LinkedList
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		linkedlist: NewLinkedList(),
	}
}

func (l *LinkedListQueue) EnQueue(e interface{}) {
	l.linkedlist.AddLast(e)
}

func (l *LinkedListQueue) DeQueue() interface{} {
	return l.linkedlist.RemoveFirst()
}

func (l *LinkedListQueue) Front() interface{} {
	return l.linkedlist.GetFirst()
}

func (l *LinkedListQueue) IsEmpty() bool {
	return l.linkedlist.IsEmpty()
}

func (l *LinkedListQueue) GetSize() int {
	return l.linkedlist.GetSize()
}
