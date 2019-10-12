/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package heap

type PriorityQueue struct {
	maxHeap *MaxHeap
}

func PriorityQueueConstructor() *PriorityQueue {
	return &PriorityQueue{
		maxHeap: NewMaxHeap(),
	}
}
func (q *PriorityQueue) GetSize() int {
	return q.maxHeap.GetSize()
}
func (q *PriorityQueue) IsEmpty() bool {
	return q.maxHeap.IsEmpty()
}
func (q *PriorityQueue) GetFront() interface{} {
	return q.maxHeap.FindMax()
}
func (q *PriorityQueue) Enqueue(e interface{}) {
	q.maxHeap.Add(e)
}
func (q *PriorityQueue) Dequeue() interface{} {
	return q.maxHeap.ExtractMax()
}
func (q *PriorityQueue) String() string {

	return q.maxHeap.data.String()
}
