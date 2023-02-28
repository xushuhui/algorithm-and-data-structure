package maps

import (
	"data-structures/datastructure/avl_tree"
)

type AvlMap struct {
	Avl *avl_tree.Avl
}

func NewAvlMap() *AvlMap {
	return &AvlMap{
		Avl: avl_tree.NewAvl(),
	}
}

func (a *AvlMap) Add() {
}

func (a *AvlMap) Remove() {
}

func (a *AvlMap) Contains() {
}

func (a *AvlMap) Get() {
}

func (a *AvlMap) Set() {
}
