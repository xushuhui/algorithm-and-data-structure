/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package segment_tree

import (
	"bytes"
	"fmt"
)

type SegmentTree struct {
	tree   []interface{}
	data   []interface{}
	merger func(interface{}, interface{}) interface{}
}

func NewST(arr []int, merger func(interface{}, interface{}) interface{}) *SegmentTree {

	data := make([]interface{}, len(arr))
	for i := 0; i < len(arr); i++ {
		data[i] = arr[i]
	}
	tree := make([]interface{}, len(arr)*4)
	st := &SegmentTree{
		data:   data,
		tree:   tree,
		merger: merger,
	}

	st.buildSegmentTree(0, 0, len(data)-1)
	return st
}
func (this *SegmentTree) buildSegmentTree(treeIndex, left, right int) {
	if left == right {
		this.tree[treeIndex] = this.data[left]
		return
	}
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)
	mid := left + (right-left)/2
	this.buildSegmentTree(leftTreeIndex, left, mid)
	this.buildSegmentTree(rightTreeIndex, mid+1, right)
	this.tree[treeIndex] = this.merger(this.tree[leftTreeIndex], this.tree[rightTreeIndex])
}
func (this *SegmentTree) GetSize() int {
	return len(this.data)
}
func (this *SegmentTree) Get(index int) interface{} {
	if index < 0 || index >= len(this.data) {
		panic("index is illegal")
	}
	return this.data[index]
}
func leftChild(index int) int {
	return 2*index + 1
}
func rightChild(index int) int {
	return 2*index + 2
}
func (this *SegmentTree) Query(queryLeft, queryRight int) interface{} {
	if queryLeft < 0 || queryLeft >= len(this.data) || queryRight < 0 || queryRight >= len(this.data) || queryLeft > queryRight {
		panic("index is illegal")
	}
	return this.query(0, 0, len(this.data)-1, queryLeft, queryRight)
}
func (this *SegmentTree) query(treeIndex, left, right, queryLeft, queryRight int) interface{} {
	if left == queryLeft && right == queryRight {
		return this.tree[treeIndex]
	}
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)
	mid := left + (right-left)/2
	if queryLeft >= mid+1 {
		return this.query(rightTreeIndex, mid+1, right, queryLeft, queryRight)
	} else if queryRight <= mid {
		return this.query(leftTreeIndex, left, mid, queryLeft, queryRight)
	}
	leftResult := this.query(leftTreeIndex, left, mid, queryLeft, mid)
	rightResult := this.query(rightTreeIndex, mid+1, right, mid+1, queryRight)
	return this.merger(leftResult, rightResult)
}
func (this *SegmentTree) Set(index int, e interface{}) {
	if index < 0 || index >= len(this.data) {
		panic("index is illegal")
	}
	this.data[index] = e
	this.set(0, 0, len(this.data)-1, index, e)
}
func (this *SegmentTree) set(treeIndex, left, right, index int, e interface{}) {
	if left == right {
		this.tree[treeIndex] = e
		return
	}
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)
	mid := left + (right-left)/2
	if index >= mid+1 {
		this.set(rightTreeIndex, mid+1, right, index, e)
	} else {
		this.set(leftTreeIndex, left, mid, index, e)
	}
	this.tree[treeIndex] = this.merger(this.tree[leftTreeIndex], this.tree[rightTreeIndex])
}

func (this *SegmentTree) String() string {
	buffer := bytes.Buffer{}

	buffer.WriteString("[")
	for i := 0; i < len(this.tree); i++ {
		if this.tree[i] != nil {
			buffer.WriteString(fmt.Sprint(this.tree[i]))
		} else {
			buffer.WriteString("nil")
		}

		if i != len(this.tree)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
