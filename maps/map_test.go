/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package maps

import "testing"

func TestBSTMap(t *testing.T) {

}
func TestLinkedListMap(t *testing.T) {
	lmap := LinkedListMapConstructor()
	lmap.Add("k", "v")
	lmap.Add("k1", "v1")
	t.Log(lmap.Get("k"))
	t.Log(lmap.GetSize())
	lmap.Set("k", "Val")
	lmap.Remove("k")
	t.Log(lmap)
	t.Log(lmap.GetSize())
}
