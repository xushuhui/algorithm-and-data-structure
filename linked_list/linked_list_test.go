/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package linked_list

import "testing"

func TestLinkedList(t *testing.T) {
	linkedlist := LinkedListConstructor()
	for i := 0; i < 10; i++ {
		linkedlist.AddFirst(i)
	}
	t.Log(linkedlist)
	t.Log(linkedlist.GetSize())
}
