/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package maps

type Map interface {
	Add(key, value interface{})
	Contains(key interface{})
	Get(key interface{})
	Set(key, value interface{})
	Remove(key interface{})
	GetSize()
	IsEmpty()
}
