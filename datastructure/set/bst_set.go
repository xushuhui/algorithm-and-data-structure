package set

import "data-structures/datastructure/binary_search_tree"

type BSTSet struct {
	bst *binary_search_tree.BST
}

func NewBSTSet() *BSTSet {
	return &BSTSet{
		bst: binary_search_tree.NewBST(),
	}
}

func (b *BSTSet) Add(e interface{}) {
}

func (b *BSTSet) Contains(e interface{}) bool {
	return false
}

func (b *BSTSet) Remove(e interface{}) {
}

func (b *BSTSet) GetSize() int {
	return b.bst.GetSize()
}

func (b *BSTSet) IsEmpty() bool {
	return b.bst.IsEmpty()
}
