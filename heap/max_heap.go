/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package heap

import (
	"data-structures/array"
	"data-structures/utils"
)

type MaxHeap struct {
	data *array.Array
}

func (this *MaxHeap) GetSize() int {
	return this.data.GetSize()
}
func (this *MaxHeap) IsEmpty() bool {
	return this.data.IsEmpty()
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func parent(index int) int {
	if index == 0 {
		panic("index is illegal")
	}
	return (index - 1) / 2
}
func (this *MaxHeap) LeftChild(index int) int {
	return index*2 + 1
}
func (this *MaxHeap) RightChild(index int) int {
	return index*2 + 2
}
func (this *MaxHeap) Add(e interface{}) {
	this.data.AddLast(e)
	this.siftUp(this.data.GetSize() - 1)
}
func (this *MaxHeap) FindMax() interface{} {
	if this.data.GetSize() == 0 {
		panic("heap is empty")
	}
	return this.data.Get(0)
}
func (this *MaxHeap) ExtractMax() {

}
func (this *MaxHeap) Replace() {

}
func (this *MaxHeap) siftUp(k int) {
	for k > 0 && utils.Compare(this.data.Get(k), this.data.Get(parent(k))) > 0 {
		this.data.Swap(k, parent(k))
		k = parent(k)
	}
}
