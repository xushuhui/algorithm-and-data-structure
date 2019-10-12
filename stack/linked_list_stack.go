/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package stack

import (
	"github.com/xushuhui/data-structures/linked_list"
)

type LinkedListStack struct {
	linkedList *linked_list.LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		linkedList: linked_list.LinkedListConstructor(),
	}
}
func (this *LinkedListStack) GetSize() int {
	return this.linkedList.GetSize()
}
func (this *LinkedListStack) IsEmpty() bool {
	return this.linkedList.IsEmpty()
}
func (this *LinkedListStack) Push(e int) {
	this.linkedList.AddLast(e)
}
func (this *LinkedListStack) Pop() interface{} {
	return this.linkedList.RemoveLast()
}
func (this *LinkedListStack) Peek() interface{} {
	return this.linkedList.GetLast()
}
