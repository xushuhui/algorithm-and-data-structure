/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package heap

import (
	"github.com/xushuhui/data-structures/array"
	"github.com/xushuhui/data-structures/utils"
)

type MaxHeap struct {
	data *array.Array
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		array.Constructor(10),
	}
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
func leftChild(index int) int {
	return index*2 + 1
}
func rightChild(index int) int {
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
func (this *MaxHeap) ExtractMax() interface{} {
	ret := this.FindMax()
	this.data.Swap(0, this.data.GetSize()-1)
	this.data.RemovLast()
	this.siftDown(0)
	return ret
}
func (this *MaxHeap) Replace(e interface{}) interface{} {
	ret := this.FindMax()
	this.data.Set(0, e)
	this.siftDown(0)
	return ret
}
func (this *MaxHeap) siftDown(k int) {
	for leftChild(k) < this.data.GetSize() {
		j := leftChild(k)
		if j+1 < this.data.GetSize() && utils.Compare(this.data.Get(j+1), this.data.Get(j)) > 0 {
			j++
		}
		if utils.Compare(this.data.Get(k), this.data.Get(j)) >= 0 {
			break
		}
		this.data.Swap(k, j)
		k = j
	}
}
func (this *MaxHeap) siftUp(k int) {
	for k > 0 && utils.Compare(this.data.Get(k), this.data.Get(parent(k))) > 0 {
		this.data.Swap(k, parent(k))
		k = parent(k)
	}
}
