package set

import "testing"

func TestLinkedListSet(t *testing.T) {
	//lset := NewLinkedListSet()
	//
	//for i := 0; i < 10; i++ {
	//	lset.Add(i)
	//}
	//t.Log(lset)
	//lset.Remove(9)
	//t.Log(lset)
	//t.Log(lset.GetSize())

	aset := NewAVLSet()
	for i := 0; i < 10; i++ {
		aset.Add(i)
	}
	//t.Log(bset)
	aset.Remove(9)
	t.Log(aset)
	t.Log(aset.GetSize())
}
