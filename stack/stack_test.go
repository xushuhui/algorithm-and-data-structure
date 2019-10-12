/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package stack

import "testing"

func TestStack(t *testing.T) {
	ast := NewArrayStack(10)

	for i := 0; i < 5; i++ {
		ast.Push(i)
	}
	t.Log(ast.array)
	t.Log(ast.Pop())
	t.Log(ast.Peek())

	lst := NewLinkedListStack()
	for i := 0; i < 5; i++ {
		lst.Push(i)
	}
	t.Log(lst.linkedList)
	t.Log(lst.Pop())
	t.Log(lst.Peek())
}
