package datastructure

type AvlSet struct {
	Avl *Avl
}

func NewAvlSet() *AvlSet {
	return &AvlSet{
		Avl: NewAvl(),
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
