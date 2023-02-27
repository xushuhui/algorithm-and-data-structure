package datastructure

import (
	"bytes"
	"fmt"
)

type doublelinkedlistnode struct {
	element interface{}
	next    *doublelinkedlistnode
	prev    *doublelinkedlistnode
}

func (n *doublelinkedlistnode) String() string {
	return fmt.Sprint(n.element)
}

type DoubleLinkedList struct {
	Size      int
	dummyHead *doublelinkedlistnode
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		dummyHead: &doublelinkedlistnode{},
	}
}
func NewDoubleLinedListNode(element interface{}, perv, next *doublelinkedlistnode) *doublelinkedlistnode {
	return &doublelinkedlistnode{
		element: element,
		prev:    perv,
		next:    next,
	}
}
func (d *DoubleLinkedList) Add(index int, element interface{}) {
	if index < 0 || index > d.Size {
		panic("invalid index")
	}
	prev := d.dummyHead
	for index := 0; index < d.Size; index++ {
		prev = prev.next
	}
	prev.next = NewDoubleLinedListNode(element, prev.prev, prev.next)
	d.Size++
}
func (d *DoubleLinkedList) String() string {
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
func (d *DoubleLinkedList) Remove() {

}
func (d *DoubleLinkedList) Contains() {

}
func (d *DoubleLinkedList) RemoveElement() {

}
