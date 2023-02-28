package linkedlist

type LinkedList interface {
	Add(index int, element interface{})
	AddFirst(element interface{})
	AddLast(element interface{})
	Remove(index int) (element interface{})
	RemoveFirst() (element interface{})
	RemoveLast() (element interface{})
	RemoveElement(element interface{})
	Get(index int) (element interface{})
	GetLast() (element interface{})
	GetFirst() (element interface{})
	GetSize() int
	Set(index int, element interface{})
	Contains(element interface{}) bool
	IsEmpty() bool
}
