/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package linked_list

import "fmt"

type Node struct {
	e    interface{}
	node *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

type LinkedList struct {
	head *Node
	size int
}

func (this *LinkedList) GetSize() {

}
func (this *LinkedList) IsEmpty() {

}
func (this *LinkedList) AddFirst() {

}
func (this *LinkedList) Add() {

}
func (this *LinkedList) AddLast() {

}
func (this *LinkedList) Remove() {

}
func (this *LinkedList) RemoveLast() {

}
func (this *LinkedList) RemoveFirst() {

}
func (this *LinkedList) Get() {

}
func (this *LinkedList) GetFirst() {

}
func (this *LinkedList) GetLast() {

}
func (this *LinkedList) Contains() {

}
func (this *LinkedList) Set() {

}
