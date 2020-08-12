/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package union_find

type UF interface {
	GetSize() int
	IsConnected(p, q int) bool
	UnionElements(p, q int)
}
