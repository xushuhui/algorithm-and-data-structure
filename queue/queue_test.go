/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package queue

import (
	"testing"
	"time"
)

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(10)
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	t.Log(q.array)
	ret := q.Dequeue()
	t.Log(ret)
	t.Log(q.GetFront())
	t.Log(q.array)
}
func TestLoopQueue(t *testing.T) {
	q := NewLoopQueue(10)
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	t.Log(q.data)
	t.Log(q.Dequeue())
	t.Log(q.GetFront())
	t.Log(q.data)
}
func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue()
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	t.Log(q.linked_list)
	t.Log(q.Dequeue())
	t.Log(q.GetFront())
	t.Log(q.linked_list)
}
func TestCompareQueue(t *testing.T) {
	n := 100000
	//s1 := utils.GenerateRandomArray(n, 0, n)
	lq := NewLoopQueue(n)
	start1 := time.Now()
	for i := 0; i < n; i++ {
		lq.Enqueue(i)
	}
	for i := 0; i < n; i++ {
		lq.Dequeue()
	}
	t.Log("LoopQueue time spent:", time.Since(start1).Seconds())
	aq := NewArrayQueue(n)
	//var s2 = make([]int, n)
	//copy(s2, s1)
	start2 := time.Now()
	for i := 0; i < n; i++ {
		aq.Enqueue(i)
	}
	for i := 0; i < n; i++ {
		aq.Dequeue()
	}
	t.Log("ArrayQueue time spent:", time.Since(start2).Seconds())
}
