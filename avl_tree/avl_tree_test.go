/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package avl_tree

import "testing"

func TestAVL(t *testing.T) {
	avl := NewAVLTree()
	avl.Add(1, 1)
	avl.Add(2, 2)
	avl.Add(3, 3)
	t.Log(avl)
	t.Log(avl.IsBalanced())
	//avl2 := NewAVLTree()
	//avl2.Add(2,2)
	//avl2.Add(1,1)
	//avl2.Add(3,3)
	//t.Log(avl2.IsBST())
	//t.Log(avl2.IsBalanced())
}
