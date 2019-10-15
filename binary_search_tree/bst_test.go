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
	bst1 := NewBST()
	bst2 := NewBST()
	bst1.root = &Node{1,
		&Node{3,
			&Node{7,
				&Node{9, nil, nil}, nil}, nil},
		&Node{5, nil, nil}}
	bst2.root = &Node{2,
		&Node{4,
			&Node{8,
				&Node{10, nil, nil}, nil}, nil},
		&Node{6, nil, nil}}
	t.Log(bst1)
}
func Sum(bst1, bst2 *BST) {
	sum(bst1.root)
}
func sum(node *Node) {

}
