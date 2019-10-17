/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package binary_search_tree

import (
	"testing"
)

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
	keys1 := bst1.BSTSlice()
	keys2 := bst2.BSTSlice()
	var sum int
	for i := 0; i < 5; i++ {
		sum += keys1[i].(int) + keys2[i].(int)
	}
	t.Log(sum)
}
func TestBST2(t *testing.T) {
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
	t.Log(Sum(bst1, bst2))
}
func inOrderBST2(node1, node2 *Node, sum *int) {
	if node1 == nil {
		return
	}
	inOrderBST2(node1.left, node2.left, sum)
	*sum += node1.e.(int) + node2.e.(int)
	inOrderBST2(node1.right, node2.right, sum)
}
func Sum(bst1, bst2 *BST) int {
	var sum int
	inOrderBST2(bst1.root, bst2.root, &sum)
	return sum
}
