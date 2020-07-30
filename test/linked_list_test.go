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

func TestLinkedList(t *testing.T) {
	l := datastructure.NewLinkedList()
	l.AddFirst(1)

	t.Log(l)
}
