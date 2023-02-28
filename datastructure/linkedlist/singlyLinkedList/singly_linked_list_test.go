package singlyLinkedList

import (
	"testing"

	tis "github.com/matryer/is"
)

func TestLinkedList_Add(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.Add(0, "1")
	is.Equal(l.Get(0), "1")
}

func TestLinkedList_AddFirst(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddFirst("1")
	is.Equal(l.Get(0), "1")
}

func TestLinkedList_AddLast(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddLast("1")

	is.Equal(l.GetLast(), "1")
}

func TestLinkedList_Clear(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddLast("1")
	l.Clear()
	is.Equal(l.size, 0)
	is.Equal(l.dummyHead.next.element, nil)
}

func TestLinkedList_Contains(t *testing.T) {

}

func TestLinkedList_Get(t *testing.T) {

}

func TestLinkedList_GetFirst(t *testing.T) {

}

func TestLinkedList_GetLast(t *testing.T) {

}

func TestLinkedList_GetSize(t *testing.T) {

}

func TestLinkedList_IsEmpty(t *testing.T) {

}

func TestLinkedList_Remove(t *testing.T) {

}

func TestLinkedList_RemoveDouble(t *testing.T) {

}

func TestLinkedList_RemoveElement(t *testing.T) {

}

func TestLinkedList_RemoveFirst(t *testing.T) {

}

func TestLinkedList_RemoveLast(t *testing.T) {
}

func TestLinkedList_Reverse(t *testing.T) {

}

func TestLinkedList_Set(t *testing.T) {

}

func TestNewLinkedList(t *testing.T) {

}

func TestNewNode(t *testing.T) {

}
