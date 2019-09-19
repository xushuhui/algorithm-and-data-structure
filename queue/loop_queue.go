/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package queue

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func LoopQueueConstructor(capacity int) (l *LoopQueue) {

	l = &LoopQueue{
		data: make([]interface{}, capacity+1),
	}
	l.front = 0
	l.tail = 0
	l.size = 0
	return
}
func (this *LoopQueue) GetSize() int {
	return this.size
}
func (this *LoopQueue) IsEmpty() bool {
	return this.front == this.tail
}
func (this *LoopQueue) GetCapacity() int {
	return len(this.data) - 1
}
func (this *LoopQueue) resize(capacity int) {
	newData := make([]interface{}, capacity+1)
	for i := 0; i < this.size; i++ {
		newData[i] = this.data[(i+this.front)%len(this.data)]
	}
	this.data = newData
	this.front = 0
	this.tail = this.size

}
func (this *LoopQueue) Enqueue(e int) {

	if (this.tail+1)%len(this.data) == this.front {
		this.resize(this.GetCapacity() * 2)
	}
	this.data[this.tail] = e
	this.tail = (this.tail + 1) % len(this.data)
	this.size++
}
func (this *LoopQueue) Dequeue() interface{} {
	if this.IsEmpty() {
		panic("Queue is empty")
	}

	ret := this.data[this.front]
	this.data[this.front] = nil
	this.front = (this.front + 1) % len(this.data)
	this.size--

	if this.size == this.GetCapacity()/4 && this.size/2 != 0 {
		this.resize(this.GetCapacity() / 2)
	}

	return ret
}
func (this *LoopQueue) GetFront() interface{} {
	if this.IsEmpty() {
		panic("Queue is empty")
	}
	return this.data[this.front]
}
