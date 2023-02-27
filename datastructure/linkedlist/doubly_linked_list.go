package linkedlist

import (
	"bytes"
	"fmt"
)

type doublylinkedlistnode struct {
	element interface{}
	next    *doublylinkedlistnode
	prev    *doublylinkedlistnode
}

func (n *doublylinkedlistnode) String() string {
	return fmt.Sprint(n.element)
}

type DoublyLinkedList struct {
	Size      int
	dummyHead *doublylinkedlistnode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		dummyHead: &doublylinkedlistnode{},
	}
}
func NewDoublyLinedListNode(element interface{}, perv, next *doublylinkedlistnode) *doublylinkedlistnode {
	return &doublylinkedlistnode{
		element: element,
		prev:    perv,
		next:    next,
	}
}
func (d *DoublyLinkedList) Add(index int, element interface{}) {
	if index < 0 || index > d.Size {
		panic("invalid index")
	}
	prev := d.dummyHead
	for index := 0; index < d.Size; index++ {
		prev = prev.next
	}
	prev.next = NewDoublyLinedListNode(element, prev.prev, prev.next)
	d.Size++
}
func (d *DoublyLinkedList) String() string {
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
func (d *DoublyLinkedList) Remove() {

}
func (d *DoublyLinkedList) Contains() {

}
func (d *DoublyLinkedList) RemoveElement() {

}
