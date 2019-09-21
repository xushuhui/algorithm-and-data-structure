package set

import "testing"

func TestLinkedListSet(t *testing.T) {
	lset := LinkedListSetConstructor()

	for i := 0; i < 10; i++ {
		lset.Add(i)
	}
	t.Log(lset)
	lset.Remove(9)
	t.Log(lset)
	t.Log(lset.GetSize())
}
