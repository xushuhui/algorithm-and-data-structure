/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package avl_tree

import (
	"bytes"
	"data-structures/utils"
	"fmt"
	"math"
)

type Node struct {
	key         interface{}
	value       interface{}
	left, right *Node
	height      int
}
type AVLTree struct {
	root *Node
	size int
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		root: NewNode(),
	}
}

func NewNode() *Node {
	return &Node{key: 0, value: 0, height: 1}
}
func (this *AVLTree) GetSize() int {
	return this.size
}
func (this *AVLTree) IsEmpty() bool {
	return this.size == 0
}

// 判断该二叉树是否是一棵二分搜索树
func (this *AVLTree) IsBST() bool {
	keys := make([]interface{}, 0)
	inOrder(this.root, &keys)
	fmt.Println(keys)
	for i := 1; i < len(keys); i++ {
		if utils.Compare(keys[i-1], keys[i]) > 0 {
			return false
		}
	}
	return true
}
func inOrder(node *Node, keys *[]interface{}) {
	if node == nil {
		return
	}
	inOrder(node.left, keys)
	*keys = append(*keys, node.key)
	inOrder(node.right, keys)

}

// 判断该二叉树是否是一棵平衡二叉树
func (this *AVLTree) IsBalanced() bool {
	return isBalanced(this.root)
}
func isBalanced(node *Node) bool {
	if node == nil {
		return true
	}
	balanceFactor := getBalanceFactor(node)
	if int(math.Abs(float64(balanceFactor))) > 1 {
		return false
	}
	return isBalanced(node.left) && isBalanced(node.right)
}
func getBalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return getHeight(node.left) - getHeight(node.right)
}
func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2
func rightRotate(y *Node) *Node {
	x := y.left
	t3 := x.right
	x.right = y
	y.left = t3
	y.height = int(math.Max(float64(getHeight(y.left)), float64(getHeight(y.right))))
	x.height = int(math.Max(float64(getHeight(x.left)), float64(getHeight(x.right))))
	return x
}

// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func leftRotate(y *Node) *Node {
	x := y.right
	t2 := x.left
	x.left = y
	y.right = t2
	y.height = int(math.Max(float64(getHeight(y.left)), float64(getHeight(y.right))))
	x.height = int(math.Max(float64(getHeight(x.left)), float64(getHeight(x.right))))
	return x
}
func (this *AVLTree) Add(key, value interface{}) {
	this.root = this.add(this.root, key, value)
}
func (this *AVLTree) add(node *Node, key, value interface{}) *Node {
	if node == nil {
		this.size++
		return &Node{
			key:   key,
			value: value,
		}
	}
	if utils.Compare(key, node.key) < 0 {
		node.left = this.add(node.left, key, value)
	} else if utils.Compare(key, node.key) > 0 {
		node.right = this.add(node.right, key, value)
	} else {
		node.value = value
	}
	node.height = 1 + int(math.Max(float64(getHeight(node.left)), float64(getHeight(node.right))))
	balanceFactor := getBalanceFactor(node)
	// LL
	if balanceFactor > 1 && getBalanceFactor(node.left) >= 0 {
		return rightRotate(node)
	}
	// RR
	if balanceFactor < -1 && getBalanceFactor(node.right) <= 0 {
		return leftRotate(node)
	}
	// LR
	if balanceFactor > 1 && getBalanceFactor(node.left) < 0 {
		node.left = leftRotate(node)
		return rightRotate(node)
	}
	// RL
	if balanceFactor < -1 && getBalanceFactor(node.right) > 0 {
		node.right = rightRotate(node)
		return leftRotate(node)
	}
	return node
}
func (this *AVLTree) getNode(node *Node, key interface{}) *Node {
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
func (this *AVLTree) Contains(key interface{}) bool {
	return this.getNode(this.root, key) != nil
}
func (this *AVLTree) Get(key interface{}) interface{} {
	node := this.getNode(this.root, key)
	if node != nil {
		return node.value
	}
	return node
}

func (this *AVLTree) Set(key, newValue interface{}) {
	node := this.getNode(this.root, key)
	if node == nil {
		panic(fmt.Sprint(key) + "not exist")
	}
	node.value = newValue
}
func (this *AVLTree) Remove(key interface{}) interface{} {
	node := this.getNode(this.root, key)
	if node != nil {
		this.root = this.remove(this.root, key)
		return node.value
	}
	return nil
}
func (this *AVLTree) remove(node *Node, key interface{}) *Node {
	if node == nil {
		return nil
	}
	var retNode *Node
	if utils.Compare(key, node.key) < 0 {
		node.left = this.remove(node.left, key)
		retNode = node
	} else if utils.Compare(key, node.key) > 0 {
		node.right = this.remove(node.right, key)
		retNode = node
	} else {
		// 待删除节点左子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			this.size--
			retNode = rightNode
		} else if node.right == nil { // 待删除节点右子树为空的情况
			leftNode := node.left
			node.left = nil
			this.size--
			retNode = leftNode
		} else {
			// 待删除节点左右子树均不为空的情况
			// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
			// 用这个节点顶替待删除节点的位置
			n := minimum(node.right)
			n.right = this.remove(node.right, n.key)
			n.left = node.left
			node.left = nil
			node.right = nil
			retNode = n
		}
	}
	if retNode == nil {
		return nil
	}
	retNode.height = 1 + int(math.Max(float64(getHeight(retNode.left)), float64(getHeight(retNode.right))))
	balanceFactor := getBalanceFactor(retNode)
	// LL
	if balanceFactor > 1 && getBalanceFactor(retNode.left) >= 0 {
		return rightRotate(retNode)
	}
	// RR
	if balanceFactor < -1 && getBalanceFactor(retNode.right) <= 0 {
		return leftRotate(retNode)
	}
	// LR
	if balanceFactor > 1 && getBalanceFactor(retNode.left) < 0 {
		retNode.left = leftRotate(retNode)
		return rightRotate(retNode)
	}
	// RL
	if balanceFactor < -1 && getBalanceFactor(retNode.right) > 0 {
		retNode.right = rightRotate(retNode)
		return leftRotate(retNode)
	}
	return retNode
}
func minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return minimum(node.left)
}
func (this *AVLTree) Minimum() interface{} {
	if this.size == 0 {
		panic("BST is empty")
	}
	return minimum(this.root).value
}

func (this *AVLTree) String() string {
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

	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(n.value) + "\n")
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
