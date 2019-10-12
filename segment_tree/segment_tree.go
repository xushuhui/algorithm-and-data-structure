/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package segment_tree

type SegmentTree struct {
	tree   interface{}
	data   []int
	Merger func()
}

func NewST() {

}
func buildSegmentTree(treeIndex, left, right int) {

}
func (this *SegmentTree) GetSize() int {
	return len(this.data)
}
func (this *SegmentTree) Get(index int) int {
	if index < 0 || index >= len(this.data) {
		panic("index is illegal")
	}
	return this.data[index]
}
func leftChild(index int) int {
	return 2*index + 1
}
func rightChild(index int) int {
	return 2*index + 2
}
func Query(queryL, queryR int) {

}
func query(treeIndex, left, right int) {

}
func Set(index int, e interface{}) {

}
func set(treeIndex, left, right, index int, e interface{}) {

}

//func (this *Array) String() string {
//	buffer := bytes.Buffer{}
//
//	buffer.WriteString("[")
//	for i := 0; i < len(this.tree); i++ {
//		if this.tree[i] != nil {
//			buffer.WriteString(fmt.Sprint(this.tree[i]))
//		} else {
//			buffer.WriteString("nil")
//		}
//
//		if i != len(this.tree)-1 {
//			buffer.WriteString(", ")
//		}
//	}
//	buffer.WriteString("]")
//
//	return buffer.String()
//}
