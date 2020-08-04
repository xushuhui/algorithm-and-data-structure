package datastructure

import (
	"bytes"
	"fmt"
)

type Array struct {
	data []interface{}
	size int
}

func NewArray(capacity int) *Array {
	return &Array{
		data: make([]interface{}, capacity),
		size: 0,
	}
}
func (a *Array) Add(index int, e interface{}) {
	if index < 0 || index > a.size {
		panic("invalid index")
	}
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}
func (a *Array) AddFirst(e interface{}) {
	a.Add(0, e)
}
func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("invalid index")
	}
	ret := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}

	a.size--
	a.data[a.size] = nil
	return ret
}
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}
func (a *Array) RemoveElement(e interface{}) {
	index := a.Find(e)
	if index != -1 {
		a.Remove(index)
	}
}
func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}
func (a *Array) Contains(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return true
		}
	}
	return false
}
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("invalid index")
	}
	return a.data[index]
}
func (a *Array) GetLast() interface{} {
	return a.Get(a.size - 1)
}

func (a *Array) GetFirst() interface{} {
	return a.Get(0)
}

func (a *Array) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("invalid index")
	}
	a.data[index] = e
}
func (a *Array) GetSize() int {
	return a.size
}
func (a *Array) IsEmpty() bool {
	return a.size == 0
}
func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
