package datastructure

import linkedlist "data-structures/datastructure/linkedlist"

type Set interface {
	Add(e interface{})
	Contains(e interface{}) bool
	Remove(e interface{})
	GetSize() int
	IsEmpty() bool
}
type LinkedListSet struct {
	linkedlist *linkedlist.LinkedList
}

func NewLinkedListSet() *LinkedListSet {
	return &LinkedListSet{
		linkedlist: linkedlist.NewLinkedList(),
	}
}

func (l *LinkedListSet) Add(e interface{}) {
	if !l.Contains(e) {
		l.linkedlist.AddFirst(e)
	}
}

func (l *LinkedListSet) Contains(e interface{}) bool {
	return l.linkedlist.Contains(e)
}

func (l *LinkedListSet) Remove(e interface{}) {
	l.linkedlist.RemoveElement(e)
}

func (l *LinkedListSet) GetSize() int {
	return l.linkedlist.GetSize()
}

func (l *LinkedListSet) IsEmpty() bool {
	return l.linkedlist.IsEmpty()
}
