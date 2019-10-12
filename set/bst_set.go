package set

import (
	"github.com/xushuhui/data-structures/binary_search_tree"
)

type BSTSet struct {
	BST *binary_search_tree.BST
}

func NewBSTSet() *BSTSet {
	return &BSTSet{
		BST: binary_search_tree.BSTConstructor(),
	}
}
func (this *BSTSet) Add(e interface{}) {
	if !this.Contains(e) {
		this.BST.Add(e)
	}
}
func (this *BSTSet) Remove(e interface{}) {
	this.BST.Remove(e)
}
func (this *BSTSet) Contains(e interface{}) bool {
	return this.BST.Contains(e)
}
func (this *BSTSet) GetSize() int {
	return this.BST.GetSize()
}
func (this *BSTSet) IsEmpty() bool {
	return this.BST.IsEmpty()
}
