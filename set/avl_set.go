package set

import "data-structures/avl_tree"

type AVLSet struct {
	avl_tree *avl_tree.AVLTree
}

func NewAVLSet() *AVLSet {
	return &AVLSet{
		avl_tree: avl_tree.NewAVLTree(),
	}
}

func (this *AVLSet) Add(e interface{}) {
	this.avl_tree.Add(e, nil)
}

func (this *AVLSet) Remove(e interface{}) interface{} {
	return this.avl_tree.Remove(e)
}
func (this *AVLSet) GetSize() int {
	return this.avl_tree.GetSize()
}
func (this *AVLSet) IsEmpty() bool {
	return this.avl_tree.IsEmpty()
}
func (this *AVLSet) Contains(e interface{}) bool {
	return this.avl_tree.Contains(e)
}
