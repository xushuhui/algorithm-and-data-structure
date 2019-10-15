/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package stack

import "github.com/xushuhui/data-structures/array"

type ArrayStack struct {
	array *array.Array
}

func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{
		array: array.NewArray(capacity),
	}
}

func (this *ArrayStack) GetSize() int {
	return this.array.GetSize()
}
func (this *ArrayStack) IsEmpty() bool {
	return this.array.IsEmpty()
}
func (this *ArrayStack) Push(e interface{}) {
	this.array.AddLast(e)
}
func (this *ArrayStack) Pop() interface{} {
	return this.array.RemovLast()
}
func (this *ArrayStack) Peek() interface{} {
	return this.array.GetLast()
}
