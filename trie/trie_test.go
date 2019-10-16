/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package trie

import "testing"

func TestTrie(t *testing.T) {
	tries := NewTrie()
	tries.Add("bat")
	t.Log(tries.Contains("ba"))
}
