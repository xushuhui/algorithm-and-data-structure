/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package trie

type Trie struct {
	root *Node
	size int
}
type Node struct {
	isWord bool
	next   map[string]*Node
}

func NewNode() *Node {
	return &Node{
		next: make(map[string]*Node),
	}
}
func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}
func (this *Trie) GetSize() int {
	return this.size
}
func (this *Trie) IsEmpty() bool {
	return this.size == 0
}
func (this *Trie) Add(word string) {
	cur := this.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = NewNode()
		}
		cur = cur.next[c]
	}
	if cur.isWord == false {
		cur.isWord = true
		this.size++
	}

}
func (this *Trie) Contains(word string) bool {
	cur := this.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isWord
}
func (this *Trie) IsPrefix(prefix string) bool {
	cur := this.root
	for _, w := range []rune(prefix) {
		c := string(w)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return true
}
