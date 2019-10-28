/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package segment_tree

import "testing"

func TestSegmentTree(t *testing.T) {
	nums := []int{1, 2, 3, 5}
	st := NewST(nums, func(a, b interface{}) interface{} {
		return a.(int) + b.(int)
	})
	t.Log(st)
	t.Log(st.Query(0, 2))
}
