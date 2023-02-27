package maps

import "data-structures/datastructure"

type BSTMap struct {
	BST *datastructure.BST
}

func NewBSTMap() *BSTMap {
	return &BSTMap{
		BST: datastructure.NewBST(),
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
