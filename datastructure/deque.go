package datastructure

type Deque struct {
	size int
	next *DequeNode
}
type DequeNode struct{}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque) GetSize() int {
	return d.size
}

func (d *Deque) EnQueue(e interface{}) {
}

func (d *Deque) DeQueue() interface{} {
	return nil
}

func (d *Deque) Front() interface{} {
	return nil
}
