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
	this.addNode(this.root, e)
}
func (this *BST) addNode(node *Node, e interface{}) *Node {
	if node == nil {
		this.size++
		return generateNode(e)
	}
	if utils.Compare(e, node.e) < 0 {
		node.left = this.addNode(node.left, e)
	}
	if utils.Compare(e, node.e) > 0 {
		node.right = this.addNode(node.right, e)
	}
	return node
}
func (this *BST) Contains(e interface{}) {
	this.containsNode(this.root, e)
}
func (this *BST) containsNode(node *Node, e interface{}) bool {
	if node == nil {
		return false
	}
	if utils.Compare(e, node.e) < 0 {
		return this.containsNode(node.left, e)
	} else if utils.Compare(e, node.e) > 0 {
		return this.containsNode(node.right, e)
	}
	return true
}
