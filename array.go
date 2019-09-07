/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package main

type Array struct {
	data []int
	size int
}

func Constructor(capacity int) *Array {
	return &Array{
		data: make([]int, capacity),
	}
}
func (this *Array) GetSize() int {
	return this.size
}
func (this *Array) IsEmpty() bool {
	return this.size == 0
}

func (this *Array) GetCapacity() int {
	return len(this.data)
}
func (this *Array) add(index, e int) {

}
