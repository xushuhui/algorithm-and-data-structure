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

func TestArray(t *testing.T) {
	a := NewArray(2)
	a.Add(0, 1)
	t.Log(a)
	a.Add(1, 2)
	t.Log(a)
}
