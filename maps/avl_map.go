package maps

import "data-structures/avl_tree"

type AVLMap struct {
	avl_tree *avl_tree.AVLTree
}

func NewAVLMap() *AVLMap {
	return &AVLMap{
		avl_tree: avl_tree.NewAVLTree(),
	}
}
func (this *AVLMap) Add(key, value interface{}) {
	this.avl_tree.Add(key, value)
}
func (this *AVLMap) Contains(key interface{}) bool {
	return this.avl_tree.Contains(key)
}
func (this *AVLMap) Get(key interface{}) interface{} {
	return this.avl_tree.Get(key)
}
func (this *AVLMap) Set(key, value interface{}) {
	this.avl_tree.Set(key, value)
}
func (this *AVLMap) Remove(key interface{}) interface{}{
	return this.avl_tree.Remove(key)
}
func (this *AVLMap) GetSize() int {
	return this.avl_tree.GetSize()
}
func (this *AVLMap) IsEmpty() bool {
	return this.avl_tree.IsEmpty()
}
