package datastructure

import (
	"bytes"
	"fmt"
)

type LinkedList struct {
	dummyHead *LinkedListNode
	size      int
}
type LinkedListNode struct {
	e    interface{}
	next *LinkedListNode
}

func (n *LinkedListNode) String() string {
	return fmt.Sprint(n.e)
}
func NewLinkedList() *LinkedList {
	return &LinkedList{
		dummyHead: &LinkedListNode{},
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
	return retNode.e
}
func (l *LinkedList) RemoveTest(index int) {
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
func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}
func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}
func (l *LinkedList) RemoveElement(e interface{}) {
	prev := l.dummyHead
	for prev.next != nil {
		if prev.next.e == e {
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
	return current.e
}
func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}
func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}
func (l *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= l.GetSize() {
		panic("invalid index")
	}
	current := l.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.e = e
}
func (l *LinkedList) Contains(e interface{}) bool {
	current := l.dummyHead
	for current.next != nil {
		if current.e == e {
			return true
		}
		current = current.next
	}
	return false
}
func (l *LinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := l.dummyHead.next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.e) + "->")
		cur = cur.next
	}

	buffer.WriteString("NULL")

	return buffer.String()
}
