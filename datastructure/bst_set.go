package datastructure

type BSTSet struct {
	bst *BST
}

func NewBSTSet() *BSTSet {
	return &BSTSet{
		bst: NewBST(),
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
