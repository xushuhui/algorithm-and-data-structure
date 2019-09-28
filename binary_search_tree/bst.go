/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package binary_search_tree

import "data-structures/utils"

type Node struct {
	e     interface{}
	left  *Node
	right *Node
}
type BST struct {
	root *Node
	size int
}

func generateNode(e interface{}) *Node {
	return &Node{e: e}
}
func BSTConstructor() *BST {
	return &BST{}
}
func (this *BST) GetSize() int {
	return this.size
}
func (this *BST) IsEmpty() bool {
	return this.size == 0
}
func (this *BST) Add(e interface{}) {
	this.add(this.root, e)
}
func (this *BST) add(node *Node, e interface{}) *Node {
	if node == nil {
		this.size++
		return generateNode(e)
	}
	if utils.Compare(e, node.e) < 0 {
		node.left = this.add(node.left, e)
	}
	if utils.Compare(e, node.e) > 0 {
		node.right = this.add(node.right, e)
	}
	return node
}
func (this *BST) Contains(e interface{}) {
	this.contains(this.root, e)
}
func (this *BST) contains(node *Node, e interface{}) bool {
	if node == nil {
		return false
	}
	if utils.Compare(e, node.e) < 0 {
		return this.contains(node.left, e)
	} else if utils.Compare(e, node.e) > 0 {
		return this.contains(node.right, e)
	}
	return true
}
