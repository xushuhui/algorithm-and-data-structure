/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package binary_search_tree

import "testing"

func TestBST(t *testing.T) {
	bst := BSTConstructor()
	for i := 0; i < 10; i++ {
		bst.Add(i)
	}
	bst.LevelOrder2()
}
