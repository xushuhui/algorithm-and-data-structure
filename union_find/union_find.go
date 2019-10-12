/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package union_find

type UnionFind interface {
	GetSize()
	isConnected(p, q int)
	UnionElements(p, q int)
}
