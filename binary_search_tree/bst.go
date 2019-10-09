/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package binary_search_tree

import (
	"bytes"
	"data-structures/utils"
	"fmt"
)

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
	this.root = this.add(this.root, e)
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
func (this *BST) String() string {
	var buffer bytes.Buffer
	generateBSTSting(this.root, 0, &buffer)
	return buffer.String()
}

// 生成以 Node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(n *Node, depth int, buffer *bytes.Buffer) {
	if n == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(n.e) + "\n")
	generateBSTSting(n.left, depth+1, buffer)
	generateBSTSting(n.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
func (this *BST) PreOrder() {
	preOrder(this.root)
}
func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.e)
	preOrder(node.left)
	preOrder(node.right)
}
func (this *BST) InOrder() {
	inOrder(this.root)
}
func inOrder(node *Node) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Println(node.e)
	inOrder(node.right)
}
func (this *BST) PostOrder() {
	postOrder(this.root)
}
func postOrder(node *Node) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Println(node.e)
}
