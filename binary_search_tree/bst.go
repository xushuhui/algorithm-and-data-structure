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
func (this *BST) Contains(e interface{}) bool {
	return this.contains(this.root, e)
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
func (this *BST) Minimum() interface{} {
	if this.size == 0 {
		panic("BST is empty")
	}
	return minimum(this.root).e
}
func minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return minimum(node.left)
}
func (this *BST) Maximum() interface{} {
	if this.size == 0 {
		panic("BST is empty")
	}
	return maximum(this.root).e
}
func maximum(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return minimum(node.right)
}

// 从二分搜索树中删除最小值所在节点, 返回最小值
func (this *BST) RemoveMin() interface{} {
	ret := this.Minimum()
	this.root = this.removeMin(this.root)
	return ret
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *BST) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		this.size--
		return rightNode
	}
	node.left = this.removeMin(node.left)
	return node
}

// 从二分搜索树中删除最大值所在节点
func (this *BST) RemoveMax() interface{} {
	ret := this.Maximum()
	this.root = this.removeMax(this.root)
	return ret
}

// 删除掉以node为根的二分搜索树中的最大节点
// 返回删除节点后新的二分搜索树的根
func (this *BST) removeMax(node *Node) *Node {
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		this.size--
		return leftNode
	}
	node.right = this.removeMax(node.right)
	return node
}

// 从二分搜索树中删除元素为e的节点
func (this *BST) Remove(e interface{}) {
	this.root = this.remove(this.root, e)
}

// 删除掉以node为根的二分搜索树中值为e的节点, 递归算法
// 返回删除节点后新的二分搜索树的根
func (this *BST) remove(node *Node, e interface{}) *Node {
	if node == nil {
		return nil
	}
	if utils.Compare(e, node.e) < 0 {
		node.left = this.remove(node.left, e)
		return node
	} else if utils.Compare(e, node.e) > 0 {
		node.right = this.remove(node.right, e)
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
