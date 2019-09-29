/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package heap

import "data-structures/array"

type MinHeap struct {
	data *array.Array
}
func (this *MinHeap) GetSize() int {
	return this.data.GetSize()
}
func (this *MinHeap) IsEmpty() bool {
	return this.data.IsEmpty()
}