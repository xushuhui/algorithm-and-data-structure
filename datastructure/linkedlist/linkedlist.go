package linkedlist

type LinkedList interface {
	Add(index int, element interface{}) (err error)
	AddFirst(element interface{}) (err error)
	AddLast(element interface{}) (err error)
	Remove(index int) (element interface{}, err error)
	RemoveFirst() (element interface{}, err error)
	RemoveLast() (element interface{}, err error)
	RemoveElement(element interface{}) (err error)
	Get(index int) (element interface{}, err error)
	GetLast() (element interface{}, err error)
	GetFirst() (element interface{}, err error)
	GetSize() int
	Set(index int, element interface{}) (err error)
	Contains(element interface{}) bool
	IsEmpty() bool
}
