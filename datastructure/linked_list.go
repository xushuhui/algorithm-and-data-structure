package datastructure

type LinkedList struct {
	dummyHead *LinkedListNode
	size      int
}
type LinkedListNode struct {
	e    interface{}
	next *LinkedListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		dummyHead: nil,
	}
}

func NewLinkedListNode(e interface{}, next *LinkedListNode) *LinkedListNode {
	return &LinkedListNode{
		e:    e,
		next: next,
	}
}
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.GetSize() {
		panic("invalid index")
	}
	prev := l.dummyHead
	//全部往前移动
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = NewLinkedListNode(e, prev.next)
	l.size++
}
func (l *LinkedList) GetSize() int {
	return l.size
}
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}
func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size-1, e)
}
