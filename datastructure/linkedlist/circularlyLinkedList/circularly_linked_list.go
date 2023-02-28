package circularlyLinkedList

import "fmt"

type node struct {
	element interface{}
	next    *node
}

func (n *node) String() string {
	return fmt.Sprint(n.element)
}

type LinkedList struct {
	size      int
	dummyHead *node
}

func (l *LinkedList) Add(index int, element interface{}) {
	panic("implement me")
}

func (l *LinkedList) AddLast(element interface{}) {
	l.Add(l.size, element)
}

func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) Remove(index int) (element interface{}) {
	panic("implement me")
}

func (l *LinkedList) RemoveElement(element interface{}) {
	panic("implement me")
}

func (l *LinkedList) Get(index int) (element interface{}) {
	panic("implement me")
}

func (l *LinkedList) Set(index int, element interface{}) {
	panic("implement me")
}

func (l *LinkedList) Contains(element interface{}) bool {
	panic("implement me")
}
