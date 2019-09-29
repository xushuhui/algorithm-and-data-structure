/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package array

import (
	"testing"
)

func TestArray(t *testing.T) {
	arr := Constructor(20)

	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}

	arr.RemovLast()

	t.Log(arr)

}
