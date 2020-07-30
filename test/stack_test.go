/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package test

import (
	"data-structures/datastructure"
	"testing"
)

func TestStack(t *testing.T) {
	l := datastructure.NewLinkedListStack()
	for i := 0; i < 5; i++ {
		l.Push(i)
	}

	t.Log(l.Print())
	t.Log(l.Pop())
	t.Log(l.Print())
	t.Log(l.Peek())
}
