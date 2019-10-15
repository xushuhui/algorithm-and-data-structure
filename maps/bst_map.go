/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package maps

import (
	"data-structures/utils"
	"fmt"
)

type BSTMap struct {
	root *BSTNode
	size int
}
type BSTNode struct {
	key         interface{}
	value       interface{}
	left, right *BSTNode
}

func NewBSTNode(key, value interface{}) *BSTNode {
	return &BSTNode{key: key, value: value}
}

// BSTMapConstructor
func NewBSTMap() *BSTMap {
	return &BSTMap{nil, 0}
}
func (this *BSTMap) GetSize() int {
	return this.size
}
func (this *BSTMap) IsEmpty() bool {
	return this.size == 0
}
func (this *BSTMap) Add(key, value interface{}) {
	this.root = this.add(this.root, key, value)
}
func (this *BSTMap) add(node *BSTNode, key, value interface{}) *BSTNode {
	if node == nil {
		this.size++
		return NewBSTNode(key, value)
	}
	if utils.Compare(key, node.key) < 0 {
		node.left = this.add(node.left, key, value)
	} else if utils.Compare(key, node.key) > 0 {
		node.right = this.add(node.right, key, value)
	} else {
		node.value = value
	}
	return node
}
func (this *BSTMap) getNode(node *BSTNode, key interface{}) *BSTNode {
	if node == nil {
		return nil
	}
	if utils.Compare(key, node.key) < 0 {
		return this.getNode(node.left, key)
	} else if utils.Compare(key, node.key) > 0 {
		return this.getNode(node.right, key)
	} else {
		return node
	}

}
func (this *BSTMap) Contains(key interface{}) bool {
	return this.getNode(this.root, key) != nil
}
func (this *BSTMap) Get(key interface{}) interface{} {
	node := this.getNode(this.root, key)
	if node != nil {
		return node.value
	}
	return node
}

func (this *BSTMap) Set(key, newValue interface{}) {
	node := this.getNode(this.root, key)
	if node == nil {
		panic(fmt.Sprint(key) + "not exist")
	}
	node.value = newValue
}
func (this *BSTMap) Remove(key interface{}) interface{} {
	node := this.getNode(this.root, key)
	if node != nil {
		this.root = this.remove(this.root, key)
		return node.value
	}
	return nil
}
func (this *BSTMap) remove(node *BSTNode, key interface{}) *BSTNode {
	if node == nil {
		return nil
	}
	if utils.Compare(key, node.key) < 0 {
		node.left = this.remove(node.left, key)
		return node
	} else if utils.Compare(key, node.key) > 0 {
		node.right = this.remove(node.right, key)
		return node
	} else {
		// 待删除节点左子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			this.size--
			return rightNode
		}
		// 待删除节点右子树为空的情况
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			this.size--
			return leftNode
		}
		// 待删除节点左右子树均不为空的情况
		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		n := minimum(node.right)
		n.right = this.removeMin(node.right)
		n.left = node.left
		node.left = nil
		node.right = nil
		return n
	}
}
func minimum(node *BSTNode) *BSTNode {
	if node.left == nil {
		return node
	}
	return minimum(node.left)
}
func (this *BSTMap) Minimum() interface{} {
	if this.size == 0 {
		panic("BST is empty")
	}
	return minimum(this.root).value
}

// 从二分搜索树中删除最小值所在节点, 返回最小值
func (this *BSTMap) RemoveMin() interface{} {
	ret := this.Minimum()
	this.root = this.removeMin(this.root)
	return ret
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *BSTMap) removeMin(node *BSTNode) *BSTNode {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		this.size--
		return rightNode
	}
	node.left = this.removeMin(node.left)
	return node
}
