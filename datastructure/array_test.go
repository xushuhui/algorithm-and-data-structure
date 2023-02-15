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
	a := NewArray(8)
	for i := 0; i < 5; i++ {
		a.Add(i,i)
	}

	t.Log(a)
	a.AddV2(1, 2)
	t.Log(a)

}
