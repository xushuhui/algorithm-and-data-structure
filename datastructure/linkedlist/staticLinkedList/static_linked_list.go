package staticLinkedList

import "fmt"

type node struct {
	element interface{}
	next    *node
}

func (n *node) String() string {
	return fmt.Sprint(n.element)
}

type LinkedList struct {
	Size      int
	dummyHead *node
}

func (l LinkedList) Add(index int, element interface{}) {
	panic("implement me")
}

func (l LinkedList) AddFirst(element interface{}) {
	panic("implement me")
}

func (l LinkedList) AddLast(element interface{}) {
	panic("implement me")
}

func (l LinkedList) Remove(index int) (element interface{}) {
	panic("implement me")
}

func (l LinkedList) RemoveFirst() (element interface{}) {
	panic("implement me")
}

func (l LinkedList) RemoveLast() (element interface{}) {
	panic("implement me")
}

func (l LinkedList) RemoveElement(element interface{}) {
	panic("implement me")
}

func (l LinkedList) Get(index int) (element interface{}) {
	panic("implement me")
}

func (l LinkedList) GetLast() (element interface{}) {
	panic("implement me")
}

func (l LinkedList) GetFirst() (element interface{}) {
	panic("implement me")
}

func (l LinkedList) GetSize() int {
	panic("implement me")
}

func (l LinkedList) Set(index int, element interface{}) {
	panic("implement me")
}

func (l LinkedList) Contains(element interface{}) bool {
	panic("implement me")
}

func (l LinkedList) IsEmpty() bool {
	panic("implement me")
}
