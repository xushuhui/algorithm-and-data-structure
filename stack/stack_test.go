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
	st := Constructor(10)

	for i := 0; i < 5; i++ {
		st.Push(i)
	}
	t.Log(st.array)
	ret := st.Pop()
	t.Log(ret)
	t.Log(st.Peek())
}
