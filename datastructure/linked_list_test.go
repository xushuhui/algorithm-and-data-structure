/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package datastructure

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.AddLast(i)
	}
	t.Log(l)
	l.RemoveTest(3)
	t.Log(l)
	//l.AddFirst(1)
	//l.AddLast(2)
	//t.Log(l)
	//l.RemoveFirst()
	//t.Log(l.GetSize())
	//val := l.Get(0)
	//t.Log(val)
	//l.Set(1,5)
	//t.Log(l)
	//l.RemoveElement(2)
	//t.Log(l)

}
func TestLinkedListMap(t *testing.T) {
	m := NewLinkedListMap()
	m.Add("k1", "v1")
	t.Log(m)
	m.Add("k2", "v2")
	t.Log(m)
	m.Add("k3", "v3")
	t.Log(m)
	m.Remove("k3")
	t.Log(m)
}
