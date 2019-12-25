/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package red_black_tree

import (
	"bytes"
	"data-structures/utils"
	"fmt"
)

const RED bool = true
const BLACK bool = false

type Node struct {
	key         interface{}
	value       interface{}
	left, right *Node
	color       bool
}
type RBTree struct {
	root *Node
	size int
}

func NewNode(key, value interface{}) *Node {
	return &Node{key: key, value: value, color: RED}
}
func NewRBTree() *RBTree {
	return &RBTree{}
}
func (t *RBTree) GetSize() int {
	return t.size
}
func (t *RBTree) IsEmpty() bool {
	return t.size == 0
}
func isRed(node *Node) bool {
	if node == nil {
		return BLACK
	}
	return node.color
}

//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func leftRotate() {

}

//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2
func rightRotate() {

}
func (t *RBTree) Add(key, value interface{}) {
	t.root = t.add(t.root, key, value)
	t.root.color = BLACK
}
func (t *RBTree) add(node *Node, key, value interface{}) *Node {
	if node == nil {
		t.size++
		return NewNode(key, value)

	}
	if utils.Compare(key, node.key) < 0 {
		node.left = t.add(node.left, key, value)
	} else if utils.Compare(key, node.key) > 0 {
		node.right = t.add(node.right, key, value)
	} else {
		node.value = value
	}
	return node
}
func (t *RBTree) getNode(node *Node, key interface{}) *Node {
	if node == nil {
		return nil
	}
	if utils.Compare(key, node.key) < 0 {
		return t.getNode(node.left, key)
	} else if utils.Compare(key, node.key) > 0 {
		return t.getNode(node.right, key)
	} else {
		return node
	}

}

func (t *RBTree) Contains(key interface{}) bool {
	return t.getNode(t.root, key) != nil
}
func (t *RBTree) Get(key interface{}) interface{} {
	node := t.getNode(t.root, key)
	if node != nil {
		return node.value
	}
	return node
}

func (t *RBTree) Set(key, newValue interface{}) {
	node := t.getNode(t.root, key)
	if node == nil {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}
	node.value = newValue
}
func (t *RBTree) Remove(key interface{}) interface{} {
	node := t.getNode(t.root, key)
	if node != nil {
		t.root = t.remove(t.root, key)
		return node.value
	}
	return nil
}
func (t *RBTree) remove(node *Node, key interface{}) *Node {
	if node == nil {
		return nil
	}
	if utils.Compare(key, node.key) < 0 {
		node.left = t.remove(node.left, key)
		return node
	} else if utils.Compare(key, node.key) > 0 {
		node.right = t.remove(node.right, key)
		return node
	} else {
		// 待删除节点左子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			t.size--
			return rightNode
		}
		// 待删除节点右子树为空的情况
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			t.size--
			return leftNode
		}
		// 待删除节点左右子树均不为空的情况
		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		n := minimum(node.right)
		n.right = t.removeMin(node.right)
		n.left = node.left
		node.left = nil
		node.right = nil
		return n
	}
}
// 返回以node为根的二分搜索树的最小值所在的节点
func minimum(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return minimum(n.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (t *RBTree) removeMin(n *Node) *Node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		t.size--

		return rightNode
	}

	n.left = t.removeMin(n.left)
	return n
}
// 获得红黑树所有的 key
func (t *RBTree) KeySet() []interface{} {
	var keySet []interface{}
	return recursive(t.root, keySet)
}

func recursive(node *Node, set []interface{}) []interface{} {
	if node == nil {
		return nil
	}
	recursive(node.left, set)
	recursive(node.right, set)
	return append(set, node.key)
}
func (t *RBTree) String() string {
	var buffer bytes.Buffer
	generateBSTSting(t.root, 0, &buffer)
	return buffer.String()
}

// 生成以node为根节点，深度为depth的描述二叉树的字符串
func generateBSTSting(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprintf("%s", node.value) + "\n")
	generateBSTSting(node.left, depth+1, buffer)
	generateBSTSting(node.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}