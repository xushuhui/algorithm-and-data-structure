/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package segment_tree

func buildSegmentTree(treeIndex,left,right int){

}
func GetSize(){

}
func Get(index int)int{

}
func leftChild(index int){

}
func rightChild(index int){
	
}
func Query(queryL,queryR int){

}
func query(treeIndex,left,right int){

}
func Set(index int,e interface{}){

}
func set(treeIndex,left,right index int,e interface{}){

}
func (this *Array) String() string {
	buffer := bytes.Buffer{}

	buffer.WriteString("[")
	for i := 0; i < len(this.tree); i++ {
		if this.tree[i] != nil {
			buffer.WriteString(fmt.Sprint(this.tree[i]))
		} else {
			buffer.WriteString("nil")
		}

		if i != len(this.tree)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
