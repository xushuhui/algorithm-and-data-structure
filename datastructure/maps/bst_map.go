package maps

import (
	"data-structures/datastructure/binary_search_tree"
)

type BSTMap struct {
	BST *binary_search_tree.BST
}

func NewBSTMap() *BSTMap {
	return &BSTMap{
		BST: binary_search_tree.NewBST(),
	}
}

func (b *BSTMap) Add() {
}

func (b *BSTMap) Remove() {
}

func (b *BSTMap) Contains() {
}

func (b *BSTMap) Get() {
}

func (b *BSTMap) Set() {
}
