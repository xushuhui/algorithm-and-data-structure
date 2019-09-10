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

func LoopQueueConstructor(capacity int) *LoopQueue {
	return &LoopQueue{
		data: make([]interface{}, capacity+1),
	}
}
func (this *LoopQueue) GetSize() {

}
func (this *LoopQueue) IsEmpty() {

}
func (this *LoopQueue) Enqueue(e int) {

}
func (this *LoopQueue) Dequeue() {

}
func (this *LoopQueue) GetFront() {

}
