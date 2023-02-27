package linkedlist

import (
	"data-structures/datastructure/maps"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.AddLast(i)
	}
	t.Log(l)
	l.Reverse()
	t.Log(l)

}

func TestDoubleLinkedList(t *testing.T) {
	l := NewDoublyLinkedList()
	for i := 0; i < 5; i++ {
		l.Add(i, i)
	}
	t.Log(l)

}

func TestLinkedListMap(t *testing.T) {
	m := maps.NewLinkedListMap()
	m.Add("k1", "v1")
	t.Log(m)
	m.Add("k2", "v2")
	t.Log(m)
	m.Add("k3", "v3")
	t.Log(m)
	m.Remove("k3")
	t.Log(m)
}
