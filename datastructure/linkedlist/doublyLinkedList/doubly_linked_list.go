package doublyLinkedList

import (
	"bytes"
	"errors"
	"fmt"
)

type node struct {
	element interface{}
	next    *node
	prev    *node
}

func (n *node) String() string {
	return fmt.Sprint(n.element)
}

type LinkedList struct {
	size      int
	dummyHead *node
}

func (l *LinkedList) AddFirst(element interface{}) (err error) {
	return l.Add(0, element)
}

func (l *LinkedList) AddLast(element interface{}) (err error) {
	return l.Add(l.size, element)
}

func (l *LinkedList) RemoveLast() (interface{}, error) {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) RemoveFirst() (interface{}, error) {
	return l.Remove(0)
}

func (l *LinkedList) GetFirst() (interface{}, error) {
	return l.Get(0)
}

func (l *LinkedList) GetLast() (interface{}, error) {
	return l.Get(l.size - 1)
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) RemoveElement(element interface{}) error {
	panic("implement me")
}

func (l *LinkedList) Get(index int) (element interface{}, err error) {
	if index < 0 || index > l.size {
		return nil, errors.New("invalid index")
	}
	prev := l.dummyHead.next
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	return prev.element, nil

}

func (l *LinkedList) Set(index int, element interface{}) error {
	panic("implement me")
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

func NewLinkedList() *LinkedList {
	return &LinkedList{
		dummyHead: &node{},
	}
}

func NewNode(element interface{}, perv, next *node) *node {
	return &node{
		element: element,
		prev:    perv,
		next:    next,
	}
}

func (l *LinkedList) Add(index int, element interface{}) (err error) {
	if index < 0 || index > l.size {
		return errors.New("invalid index")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = NewNode(element, prev.prev, prev.next)
	l.size++
	return
}

func (l *LinkedList) String() string {
	buf := bytes.Buffer{}
	cur := l.dummyHead.next
	buf.WriteString("NULL")
	for cur != nil {
		buf.WriteString("<-" + fmt.Sprint(cur.element) + "->")
		cur = cur.next
	}
	buf.WriteString("NULL")
	return buf.String()
}

func (l *LinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index > l.size {
		return nil, errors.New("invalid index")
	}
	current := l.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	retNode := current
	current.next = retNode.next
	retNode.next = nil
	retNode.prev = nil
	l.size--
	return retNode.element, nil
}
