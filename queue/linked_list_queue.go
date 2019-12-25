package queue

import "data-structures/linked_list"

type LinkedListQueue struct {
	linked_list *linked_list.LinkedList
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		linked_list: linked_list.NewLinkedList(),
	}
}
func (t *LinkedListQueue) GetSize() int {
	return t.linked_list.GetSize()
}
func (t *LinkedListQueue) IsEmpty() bool {
	return t.linked_list.IsEmpty()
}
func (t *LinkedListQueue) Enqueue(e interface{}) {
	t.linked_list.AddLast(e)
}
func (t *LinkedListQueue) Dequeue() interface{} {
	return t.linked_list.RemoveFirst()
}
func (t *LinkedListQueue) GetFront() interface{} {
	return t.linked_list.GetFirst()
}
