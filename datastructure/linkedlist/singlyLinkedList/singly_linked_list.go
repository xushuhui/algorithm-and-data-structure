package singlyLinkedList

import (
	"bytes"
	"errors"
	"fmt"
)

type node struct {
	element interface{}
	next    *node
}

func (n *node) String() string {
	return fmt.Sprint(n.element)
}

type LinkedList struct {
	dummyHead *node
	size      int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		dummyHead: &node{},
	}
}

func NewNode(element interface{}, next *node) *node {
	return &node{
		element: element,
		next:    next,
	}
}

func (l *LinkedList) Add(index int, e interface{}) (err error) {
	if index < 0 || index > l.GetSize() {
		return errors.New("invalid index")
	}
	prev := l.dummyHead
	// 全部往前移动
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = NewNode(e, prev.next)
	l.size++
	return
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) AddFirst(e interface{}) (err error) {
	return l.Add(0, e)
}

func (l *LinkedList) AddLast(e interface{}) (err error) {
	return l.Add(l.size, e)
}

func (l *LinkedList) RemoveLast() (interface{}, error) {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) RemoveFirst() (interface{}, error) {
	return l.Remove(0)
}

func (l *LinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= l.GetSize() {

		return nil, errors.New("invalid index")
	}
	prev := l.dummyHead.next
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	retNode := prev
	prev.next = retNode.next
	retNode.next = nil
	l.size--
	return retNode.element, nil
}

func (l *LinkedList) RemoveElement(element interface{}) (err error) {
	prev := l.dummyHead.next
	for prev != nil {
		if prev.element == element {
			delNode := prev
			prev.next = delNode.next
			delNode.next = nil
			l.size--
			break
		}
		prev = prev.next
	}
	return
}

func (l *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.GetSize() {
		return nil, errors.New("invalid index")
	}
	current := l.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.element, nil
}

func (l *LinkedList) GetFirst() (interface{}, error) {
	return l.Get(0)
}

func (l *LinkedList) GetLast() (interface{}, error) {
	return l.Get(l.size - 1)
}

func (l *LinkedList) Set(index int, element interface{}) error {
	if index < 0 || index >= l.GetSize() {
		return errors.New("invalid index")
	}
	current := l.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.element = element
	return nil
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
		current.element = nil
		current = current.next
	}
	l.size = 0
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
