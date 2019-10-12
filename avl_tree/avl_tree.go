/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package avl_tree

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

func NewNode() *Node {
	return &Node{key: "", value: "", height: 1}
}
func (this *AVLTree) GetSize() int {
	return this.size
}
func (this *AVLTree) IsEmpty() bool {
	return this.size == 0
}
func (this *AVLTree) Add(key, value interface{}) {
	this.root = this.add(this.root, key, value)
}
func (this *AVLTree) add(node *Node, key, value interface{}) *Node {

}
func (this *AVLTree) Contains() {

}
func (this *AVLTree) Get() interface{} {

}
func getNode(node *Node, key interface{}) {

}
func (this *AVLTree) Set(k, newValue interface{}) {

}

// 从二分搜索树中删除键为key的节点
func (this *AVLTree) Remove() {

}
func (this *AVLTree) remove() {

}
