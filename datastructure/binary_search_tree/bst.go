package binary_search_tree

type BST struct {
	Size int
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) Add() {
}

func (b *BST) GetSize() int {
	return b.Size
}

func (b *BST) IsEmpty() bool {
	return b.Size == 0
}
