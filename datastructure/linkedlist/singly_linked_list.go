package linkedlist

import (
	"bytes"
	"fmt"
)

type linkedlistnode struct {
	element interface{}
	next    *linkedlistnode
}

func (n *linkedlistnode) String() string {
	return fmt.Sprint(n.element)
}

type LinkedList struct {
	dummyHead *linkedlistnode
	size      int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		dummyHead: &linkedlistnode{},
	}
}

func NewLinkedListNode(element interface{}, next *linkedlistnode) *linkedlistnode {
	return &linkedlistnode{
		element: element,
		next:    next,
	}
}

func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.GetSize() {
		panic("invalid index")
	}
	prev := l.dummyHead
	// 全部往前移动
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
	l.Add(l.size, e)
}

func (l *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= l.GetSize() {
		panic("invalid index")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil
	l.size--
	return retNode.element
}

func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

func (l *LinkedList) RemoveElement(element interface{}) {
	prev := l.dummyHead
	for prev.next != nil {
		if prev.next.element == element {
			delNode := prev.next
			prev.next = delNode.next
			delNode.next = nil
			l.size--
			break
		}
		prev = prev.next
	}
}

func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.GetSize() {
		panic("invalid index")
	}
	current := l.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.element
}

func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

func (l *LinkedList) Set(index int, element interface{}) {
	if index < 0 || index >= l.GetSize() {
		panic("invalid index")
	}
	current := l.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.element = element
}

func (l *LinkedList) Contains(element interface{}) bool {
	current := l.dummyHead.next
	for current != nil {
		if current.element == element {
			return true
		}
		current = current.next
	}
	return false
}
func (l *LinkedList) Clear() {
	current := l.dummyHead.next
	for current != nil {
		current = nil
		current = current.next
	}
}
func (l *LinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := l.dummyHead.next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.element) + "->")
		cur = cur.next
	}

	buffer.WriteString("NULL")

	return buffer.String()
}
func (l *LinkedList) Reverse() {

	cur := l.dummyHead.next
	var i int
	for cur != nil {
		l.AddFirst(cur.element)
		cur = cur.next
		i++
		l.Remove(i)
	}

}

func (l *LinkedList) RemoveDouble(index int) {
	if index < 0 || index >= l.GetSize() {
		panic("invalid index")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	retNode := prev.next
	retNode2 := prev.next.next
	prev.next = retNode.next.next
	retNode.next = nil
	retNode2.next = nil
	l.size = l.size - 2
}
