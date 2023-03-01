package red_black_tree

// const RED bool = true
// const BLACK bool = false

//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2

//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2

type RedBlackTree struct{}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func Add() {
}

func Remove() {
}
