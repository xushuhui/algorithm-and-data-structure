package doublyLinkedList

import (
	"bytes"
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

type linkedList struct {
	Size      int
	dummyHead *node
}

func NewLinkedList() *linkedList {
	return &linkedList{
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
func (d *linkedList) Add(index int, element interface{}) {
	if index < 0 || index > d.Size {
		panic("invalid index")
	}
	prev := d.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = NewNode(element, prev.prev, prev.next)
	d.Size++
}
func (d *linkedList) String() string {
	buf := bytes.Buffer{}
	cur := d.dummyHead.next
	buf.WriteString("NULL")
	for cur != nil {
		buf.WriteString("<-" + fmt.Sprint(cur.element) + "->")
		cur = cur.next
	}
	buf.WriteString("NULL")
	return buf.String()
}
func (d *linkedList) Remove(index int) interface{} {
	if index < 0 || index > d.Size {
		panic("invalid index")
	}
	current := d.dummyHead
	for i := 0; i < index; i++ {
		current = current.next
	}
	retNode := current
	current.next = retNode.next
	retNode.next = nil
	retNode.prev = nil
	d.Size--
	return retNode.element
}
func (d *linkedList) Contains() {

}
func (d *linkedList) RemoveElement() {

}
