package stack

import (
	s "data-structures/datastructure/linkedlist/singlyLinkedList"
)

type Stack interface {
	Push(e interface{})
	Pop() interface{}
	IsEmpty() bool
	GetSize() int
	Peek() interface{}
}
type LinkedListStack struct {
	linkedlist *s.LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		linkedlist: s.NewLinkedList(),
	}
}

func (l *LinkedListStack) Push(e interface{}) {
	l.linkedlist.AddFirst(e)
}

func (l *LinkedListStack) Pop() (interface{},error) {
	return l.linkedlist.RemoveFirst()
}

func (l *LinkedListStack) Peek() (interface{},error) {
	return l.linkedlist.GetFirst()
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.linkedlist.IsEmpty()
}

func (l *LinkedListStack) GetSize() int {
	return l.linkedlist.GetSize()
}

func (l *LinkedListStack) Print() *s.LinkedList {
	return l.linkedlist
}
