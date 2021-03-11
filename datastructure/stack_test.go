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

func TestStack(t *testing.T) {
	l := NewLinkedListStack()
	for i := 0; i < 5; i++ {
		l.Push(i)
	}

	t.Log(l.Print())
	t.Log(l.Pop())
	t.Log(l.Print())
	t.Log(l.Peek())
}
