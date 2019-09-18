package maps

import "fmt"

type Node struct {
	key   interface{}
	value interface{}
	next  *Node
}
type LinkedListMap struct {
	dummyHead *Node
	size      int
}

func LinkedListMapConstructor() *LinkedListMap {
	return &LinkedListMap{
		dummyHead: &Node{},
		size:      0,
	}
}
func (this *LinkedListMap) GetSize() int {
	return this.size
}
func (this *LinkedListMap) IsEmpty() bool {
	return this.size == 0
}

func (this *LinkedListMap) Contains(key interface{}) bool {
	return this.getNode(key) != nil
}
func (this *LinkedListMap) Get(key interface{}) interface{} {
	node := this.getNode(key)
	if node == nil {
		return nil
	} else {
		return node.value
	}
}
func (this *LinkedListMap) getNode(key interface{}) *Node {
	cur := this.dummyHead.next
	for cur != nil {
		if cur.key == key {
			return cur
		}
		cur = cur.next
	}
	return nil
}
func (this *LinkedListMap) Add(key, value interface{}) {
	node := this.getNode(key)
	if node == nil {
		newNode := &Node{
			key:   key,
			value: value,
			next:  this.dummyHead.next,
		}
		this.dummyHead.next = newNode
		this.size++
	} else {
		node.value = value
	}
}
func (this *LinkedListMap) Set(key, value interface{}) {
	node := this.getNode(key)
	if node == nil {
		panic(fmt.Sprintf("%v,not exist", key))
	}
	node.value = value
}
func (this *LinkedListMap) Remove(key interface{}) interface{} {

}
