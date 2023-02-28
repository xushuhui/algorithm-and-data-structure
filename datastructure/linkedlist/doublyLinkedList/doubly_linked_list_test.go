package doublyLinkedList

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

func TestLinkedList_Contains(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddLast("1")
	is.Equal(l.Contains("1"), true)
	is.Equal(l.Contains("2"), false)
}

func TestLinkedList_Get(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddLast("1")
	is.Equal(l.Get(0), "1")
}

func TestLinkedList_GetFirst(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddFirst("1")
	is.Equal(l.GetFirst(), "1")
}

func TestLinkedList_GetLast(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddLast("1")
	is.Equal(l.GetLast(), "1")
}

func TestLinkedList_GetSize(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.Equal(l.GetSize(), 0)
	l.AddLast("1")
	is.Equal(l.GetSize(), 1)
}

func TestLinkedList_IsEmpty(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.True(l.IsEmpty())
	l.AddLast("1")
	is.True(!l.IsEmpty())
}

func TestLinkedList_Remove(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddFirst("1")
	is.Equal(l.Remove(0), "1")
}

func TestLinkedList_RemoveElement(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddFirst("1")
	l.AddFirst("2")
	l.RemoveElement("1")
	is.Equal(l.GetFirst(), "2")
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddFirst("1")
	is.Equal(l.RemoveFirst(), "1")
}

func TestLinkedList_RemoveLast(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddLast("1")
	is.Equal(l.RemoveLast(), "1")
}

func TestLinkedList_Set(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	l.AddFirst("1")
	l.AddFirst("2")
	l.Set(1, "3")
	is.Equal(l.GetLast(), "3")
}
