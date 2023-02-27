package array

import (
	"bytes"
	"fmt"
)

type Array struct {
	data []int
	size int
}

func NewArray(capacity int) *Array {
	return &Array{
		data: make([]int, capacity),
		size: 0,
	}
}

func (a *Array) Add(index int, e int) {
	if index < 0 || index > a.size {
		panic("invalid index")
	}
	// 扩容
	if a.size == len(a.data) {
		a.resize(a.size * 2)
	}
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}

func (a *Array) AddV2(index int, e int) {
	if index < 0 || index > a.size {
		panic("invalid index")
	}
	// 扩容
	if a.size == len(a.data) {
		a.resize(a.size * 2)
	}
	a.data[a.size] = a.data[index]
	a.data[index] = e
	a.size++
}

func (a *Array) AddFirst(e int) {
	a.Add(0, e)
}

func (a *Array) AddLast(e int) {
	a.Add(a.size, e)
}

func (a *Array) Remove(index int) int {
	if index < 0 || index >= a.size {
		panic("invalid index")
	}
	ret := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}

	a.size--
	a.data[a.size] = 0
	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.resize(len(a.data))
	}
	return ret
}

func (a *Array) RemoveLast() int {
	return a.Remove(a.size - 1)
}

func (a *Array) RemoveFirst() int {
	return a.Remove(0)
}

func (a *Array) RemoveElement(e int) {
	index := a.Find(e)
	if index != -1 {
		a.Remove(index)
	}
}

func (a *Array) Find(e int) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}

func (a *Array) Contains(e int) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return true
		}
	}
	return false
}

func (a *Array) Get(index int) int {
	if index < 0 || index >= a.size {
		panic("invalid index")
	}
	return a.data[index]
}

func (a *Array) GetLast() int {
	return a.Get(a.size - 1)
}

func (a *Array) GetFirst() int {
	return a.Get(0)
}

func (a *Array) Set(index int, e int) {
	if index < 0 || index >= a.size {
		panic("invalid index")
	}
	a.data[index] = e
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) GetCapacity() int {
	return len(a.data)
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

func (a *Array) resize(newCapacity int) {
	newData := make([]int, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}

func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		// fmt.Sprint 将 int 类型转换为字符串
		buffer.WriteString(fmt.Sprint(a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
