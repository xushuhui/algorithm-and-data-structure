package avl_tree

// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2

// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//
//	  y                             x
//	/  \                          /   \
//
// T1   x      向左旋转 (y)       y     z
//
//	  / \   - - - - - - - ->   / \   / \
//	T2  z                     T1 T2 T3 T4
//	   / \
//	  T3 T4
type Avl struct {
	Size int
}

func NewAvl() *Avl {
	return &Avl{}
}

func (a *Avl) GetSize() int {
	return a.Size
}

func (a *Avl) IsEmpty() bool {
	return a.Size == 0
}
