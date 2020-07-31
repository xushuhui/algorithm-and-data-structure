package datastructure

import (
	"bytes"
	"fmt"
)

type Map interface {
	Add(key, value interface{})
	Remove(key interface{}) interface{}
	Contains(key interface{}) bool
	Get(key interface{}) interface{}
	Set(key, value interface{})
	GetSize() int
	IsEmpty() bool
}
type LinkedListMap struct {
	dummyHead *mapNode
	size      int
}
type mapNode struct {
	key, value interface{}
	next       *mapNode
}

func NewMapNode(key, value interface{}, next *mapNode) *mapNode {
	return &mapNode{
		key:   key,
		value: value,
		next:  next,
	}
}
func NewLinkedListMap() *LinkedListMap {
	return &LinkedListMap{
		dummyHead: &mapNode{},
		size:      0,
	}
}
func (l *LinkedListMap) getNode(key interface{}) *mapNode {
	current := l.dummyHead.next
	for current != nil {
		if key == current.key {
			return current
		}
		current = current.next
	}
	return nil
}
func (l *LinkedListMap) Add(key, value interface{}) {
	node := l.getNode(key)
	if node == nil {
		l.dummyHead.next = NewMapNode(key, value, l.dummyHead.next)
		l.size++
	} else {
		node.value = value
	}
}
func (l *LinkedListMap) Remove(key interface{}) interface{} {
	node := l.getNode(key)
	if node == nil {
		panic("key not find")
	}
	current := l.dummyHead
	for current.next != nil {
		if current.next.key == key {
			delNode := current.next
			current.next = delNode.next
			delNode.next = nil
			l.size--
			return delNode.value
		}
		current = current.next
	}
	return nil
}
func (l *LinkedListMap) Contains(key interface{}) bool {
	return l.getNode(key) != nil
}
func (l *LinkedListMap) Get(key interface{}) interface{} {
	node := l.getNode(key)
	if node != nil {
		return node.value
	}
	return nil
}
func (l *LinkedListMap) Set(key, newValue interface{}) {
	node := l.getNode(key)
	if node == nil {
		panic("key not find")
	}
	node.value = newValue
}
func (l *LinkedListMap) GetSize() int {
	return l.size
}
func (l *LinkedListMap) IsEmpty(key, value interface{}) bool {
	return l.size == 0
}
func (l *LinkedListMap) String() string {
	buffer := bytes.Buffer{}
	cur := l.dummyHead.next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.key) + ":" + fmt.Sprint(cur.value) + "->")
		cur = cur.next
	}

	buffer.WriteString("NULL")

	return buffer.String()
}
