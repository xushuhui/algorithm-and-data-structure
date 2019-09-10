/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package queue

import "testing"

func TestQueue(t *testing.T) {
	q := Constructor(10)

	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	t.Log(q.array)
	ret := q.Dequeue()
	t.Log(ret)
	t.Log(q.GetFront())
	t.Log(q.array)
}
