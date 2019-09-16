/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package linked_list

import (
	"bytes"
	"fmt"
)

type Node struct {
	e    interface{}
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

type LinkedList struct {
	dummyHead *Node
	size      int
}

func LinkedListConstructor() *LinkedList {
	return &LinkedList{
		dummyHead: &Node{},
	}
}
func (this *LinkedList) GetSize() int {
	return this.size
}
func (this *LinkedList) IsEmpty() bool {
	return this.size == 0
}
func (this *LinkedList) AddFirst(e interface{}) {
	this.Add(0, e)
}
func (this *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > this.size {
		panic("index is illegal")
	}
	prev := this.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = &Node{e, prev.next}
	this.size++
}
func (this *LinkedList) AddLast(e interface{}) {
	this.Add(this.size, e)
}
func (this *LinkedList) Remove(index int) interface{} {
	if index < 0 || index > this.size {
		panic("index is illegal")
	}
	prev := this.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil
	this.size--
	return retNode.e
}
func (this *LinkedList) RemoveLast() interface{} {
	return this.Remove(this.size - 1)
}
func (this *LinkedList) RemoveFirst() interface{} {
	return this.Remove(0)
}
func (this *LinkedList) Get(index int) interface{} {
	if index < 0 || index > this.size {
		panic("index is illegal")
	}
	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}
func (this *LinkedList) GetFirst() interface{} {
	return this.Get(0)
}
func (this *LinkedList) GetLast() interface{} {
	return this.Get(this.size - 1)
}
func (this *LinkedList) Contains(e interface{}) bool {
	cur := this.dummyHead.next
	for cur != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false
}
func (this *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index > this.size {
		panic("index is illegal")
	}
	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}
func (this *LinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := this.dummyHead.next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.e) + "->")
		cur = cur.next
	}

	buffer.WriteString("NULL")

	return buffer.String()
}
