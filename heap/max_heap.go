package heap

import (
	"data-structures/array"
	"reflect"
)

type MaxHeap struct {
	data *array.Array
}

func (this *MaxHeap) GetSize() int {
	return this.data.GetSize()
}
func (this *MaxHeap) IsEmpty() bool {
	return this.data.IsEmpty()
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func parent(index int) int {
	if index == 0 {
		panic("index is illegal")
	}
	return (index - 1) / 2
}
func (this *MaxHeap) LeftChild(index int) int {
	return index*2 + 1
}
func (this *MaxHeap) RightChild(index int) int {
	return index*2 + 2
}
func (this *MaxHeap) Add(e interface{}) {
	this.data.AddLast(e)
	this.siftUp(this.data.GetSize() - 1)
}
func (this *MaxHeap) FindMax() interface{} {
	if this.data.GetSize() == 0 {
		panic("heap is empty")
	}
	return this.data.Get(0)
}
func (this *MaxHeap) ExtractMax() {

}
func (this *MaxHeap) Replace() {

}
func (this *MaxHeap) siftUp(k int) {
	for k > 0 && Compare(this.data.Get(k), this.data.Get(parent(k))) > 0 {

		this.data.Swap(k, parent(k))
		k = parent(k)
	}
}
func Compare(a interface{}, b interface{}) int {
	aType := reflect.TypeOf(a).String()
	bType := reflect.TypeOf(b).String()

	if aType != bType {
		panic("cannot compare different type params")
	}

	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		} else {
			return 0
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) < b.(string) {
			return -1
		} else {
			return 0
		}
	case float64:
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) < b.(float64) {
			return -1
		} else {
			return 0
		}
	default:
		panic("unsupported type params")
	}
}
