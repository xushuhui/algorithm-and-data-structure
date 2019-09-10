/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package queue

import "data-structures/array"

type ArrayQueue struct {
	array *array.Array
}

func Constructor(capacity int) *ArrayQueue {
	return &ArrayQueue{
		array: array.Constructor(capacity),
	}
}
func (this *ArrayQueue) GetSize() int {
	return this.array.GetSize()
}
func (this *ArrayQueue) IsEmpty() bool {
	return this.array.IsEmpty()
}
func (this *ArrayQueue) Enqueue(e interface{}) {
	this.array.AddLast(e)
}
func (this *ArrayQueue) Dequeue() interface{} {
	return this.array.RemoveFirst()
}
func (this *ArrayQueue) GetFront() interface{} {
	return this.array.GetFirst()
}
