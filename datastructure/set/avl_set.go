package set

import (
	"data-structures/datastructure/avl_tree"
)

type AvlSet struct {
	Avl *avl_tree.Avl
}

func NewAvlSet() *AvlSet {
	return &AvlSet{
		Avl: avl_tree.NewAvl(),
	}
}

func (a *AvlSet) Add(e interface{}) {
}

func (a *AvlSet) Contains(e interface{}) bool {
	return false
}

func (a *AvlSet) Remove(e interface{}) {
}

func (a *AvlSet) GetSize() int {
	return a.Avl.GetSize()
}

func (a *AvlSet) IsEmpty() bool {
	return a.Avl.IsEmpty()
}
