package queue

import (
	linkedlist2 "data-structures/datastructure/linkedlist/singlyLinkedList"
)

type Queue interface {
	IsEmpty() bool
	GetSize() int
	EnQueue(e interface{})
	DeQueue() interface{}
	Front() interface{}
}
type LinkedListQueue struct {
	linkedlist *linkedlist2.LinkedList
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		linkedlist: linkedlist2.NewLinkedList(),
	}
}

func (l *LinkedListQueue) EnQueue(e interface{}) {
	l.linkedlist.AddLast(e)
}

func (l *LinkedListQueue) DeQueue() (interface{},error) {
	return l.linkedlist.RemoveFirst()
}

func (l *LinkedListQueue) Front() (interface{},error) {
	return l.linkedlist.GetFirst()
}

func (l *LinkedListQueue) IsEmpty() bool {
	return l.linkedlist.IsEmpty()
}

func (l *LinkedListQueue) GetSize() int {
	return l.linkedlist.GetSize()
}
