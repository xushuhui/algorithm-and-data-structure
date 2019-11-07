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
func (this *LinkedListQueue) GetSize() int {
	return this.linked_list.GetSize()
}
func (this *LinkedListQueue) IsEmpty() bool {
	return this.linked_list.IsEmpty()
}
func (this *LinkedListQueue) Enqueue(e interface{}) {
	this.linked_list.AddLast(e)
}
func (this *LinkedListQueue) Dequeue() interface{} {
	return this.linked_list.RemoveFirst()
}
func (this *LinkedListQueue) GetFront() interface{} {
	return this.linked_list.GetFirst()
}
