/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package maps

import (
	"bytes"
	"fmt"
)

type LinkedListNode struct {
	key   interface{}
	value interface{}
	next  *LinkedListNode
}
type LinkedListMap struct {
	dummyHead *LinkedListNode
	size      int
}

func NewLinkedListMap() *LinkedListMap {
	return &LinkedListMap{
		dummyHead: &LinkedListNode{},
		size:      0,
	}
}
func (this *LinkedListMap) GetSize() int {
	return this.size
}
func (this *LinkedListMap) IsEmpty() bool {
	return this.size == 0
}

func (this *LinkedListMap) Contains(key interface{}) bool {
	return this.getNode(key) != nil
}
func (this *LinkedListMap) Get(key interface{}) interface{} {
	node := this.getNode(key)
	if node == nil {
		return nil
	} else {
		return node.value
	}
}
func (this *LinkedListMap) getNode(key interface{}) *LinkedListNode {
	cur := this.dummyHead.next
	for cur != nil {
		if cur.key == key {
			return cur
		}
		cur = cur.next
	}
	return nil
}
func (this *LinkedListMap) Add(key, value interface{}) {
	node := this.getNode(key)
	if node == nil {
		newNode := &LinkedListNode{
			key:   key,
			value: value,
			next:  this.dummyHead.next,
		}
		this.dummyHead.next = newNode
		this.size++
	} else {
		node.value = value
	}
}
func (this *LinkedListMap) Set(key, value interface{}) {
	node := this.getNode(key)
	if node == nil {
		panic(fmt.Sprintf("%v,not exist", key))
	}
	node.value = value
}
func (this *LinkedListMap) Remove(key interface{}) interface{} {
	prev := this.dummyHead
	for prev.next != nil {
		if prev.next.key == key {
			break
		}
		prev = prev.next
	}
	if prev.next != nil {
		delNode := prev.next
		prev.next = delNode.next
		delNode.next = nil
		this.size--
		return delNode.value
	}
	return nil
}

func (this *LinkedListMap) String() string {
	buffer := bytes.Buffer{}
	cur := this.dummyHead.next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.key) + "-" + fmt.Sprint(cur.value) + "->")
		cur = cur.next
	}

	buffer.WriteString("NULL")

	return buffer.String()
}
