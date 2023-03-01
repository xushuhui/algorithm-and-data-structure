package doublyLinkedList

import (
	"testing"

	tis "github.com/matryer/is"
)

func TestLinkedList_Add(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()

	is.NoErr(l.Add(0, "1"))
	result, err := l.Get(0)
	is.NoErr(err)
	is.Equal(result, "1")
}

func TestLinkedList_AddFirst(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddFirst("1"))
	result, err := l.Get(0)
	is.NoErr(err)
	is.Equal(result, "1")
}

func TestLinkedList_AddLast(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddLast("1"))
	result, err := l.GetLast()
	is.NoErr(err)
	is.Equal(result, "1")

}

func TestLinkedList_Contains(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddLast("1"))
	is.Equal(l.Contains("1"), true)
	is.Equal(l.Contains("2"), false)
}

func TestLinkedList_Get(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddLast("1"))
	result, err := l.Get(0)
	is.NoErr(err)
	is.Equal(result, "1")
}

func TestLinkedList_GetFirst(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddFirst("1"))
	result, err := l.GetFirst()
	is.NoErr(err)
	is.Equal(result, "1")

}

func TestLinkedList_GetLast(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddLast("1"))
	result, err := l.GetLast()
	is.NoErr(err)
	is.Equal(result, "1")
}

func TestLinkedList_GetSize(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.Equal(l.GetSize(), 0)
	is.NoErr(l.AddLast("1"))
	is.Equal(l.GetSize(), 1)

}

func TestLinkedList_IsEmpty(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.True(l.IsEmpty())
	is.NoErr(l.AddLast("1"))
	is.True(!l.IsEmpty())
}

func TestLinkedList_Remove(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddFirst("1"))
	result, err := l.Remove(0)
	is.NoErr(err)
	is.Equal(result, "1")

}

func TestLinkedList_RemoveElement(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddFirst("1"))
	is.NoErr(l.AddFirst("2"))
	is.NoErr(l.RemoveElement("1"))
	result, err := l.GetFirst()
	is.NoErr(err)
	is.Equal(result, "2")

}

func TestLinkedList_RemoveFirst(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddFirst("1"))
	result, err := l.RemoveFirst()
	is.NoErr(err)
	is.Equal(result, "1")
}

func TestLinkedList_RemoveLast(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddLast("1"))
	result, err := l.RemoveLast()
	is.NoErr(err)
	is.Equal(result, "1")
}

func TestLinkedList_Set(t *testing.T) {
	is := tis.New(t)
	l := NewLinkedList()
	is.NoErr(l.AddFirst("1"))
	is.NoErr(l.AddFirst("2"))
	is.NoErr(l.Set(1, "3"))
	result, err := l.GetLast()
	is.NoErr(err)
	is.Equal(result, "3")

}
