/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package hashtable

import (
	"data-structures/red_black_tree"
	"fmt"
	"hash/fnv"
	"strconv"
)

const upperTol = 10
const lowerTol = 2

var capacityIndex = 0
var capacity = []int{53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593,
	49157, 98317, 196613, 393241, 786433, 1572869, 3145739, 6291469,
	12582917, 25165843, 50331653, 100663319, 201326611, 402653189, 805306457, 1610612741};

type HashTable struct {
	hashtable []*red_black_tree.RBTree
	M         int
	size      int
}

func NewHashTable() *HashTable {
	M := capacity[capacityIndex]
	var hashtable []*red_black_tree.RBTree
	for i := 0; i < M; i++ {
		hashtable = append(hashtable, red_black_tree.NewRBTree())
	}
	return &HashTable{hashtable, M, 0}
}
func (this *HashTable) GetSize() int {
	return this.size
}
func (this *HashTable) Remove(key interface{}) interface{}{
	var ret interface{}
	maps := this.hashtable[this.hash(key)]
	if maps.Contains(key){
		ret = maps.Remove(key)
		this.size--
		if this.size <lowerTol *this.M && capacityIndex-1 >= 0 {
			capacityIndex--
			this.resize(capacity[capacityIndex])
		}
	}
	return ret
}
func (this *HashTable) Add(key, value interface{}) {
	maps := this.hashtable[this.hash(key)]
	if maps.Contains(key) {
		maps.Set(key,value)
	} else {
		maps.Add(key,value)
		this.size++
		if this.size >= upperTol *this.M && capacityIndex+1< len(capacity) {
			capacityIndex++
			this.resize(capacity[capacityIndex])
		}
	}
}
func (this *HashTable) Contains(key interface{}) bool{
	return this.hashtable[this.hash(key)].Contains(key)
}
func (this *HashTable) Get(key interface{})interface{} {
	return this.hashtable[this.hash(key)].Get(key)
}
func (this *HashTable) Set(key, value interface{}) {
	maps := this.hashtable[this.hash(key)]
	if !maps.Contains(key) {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}
	maps.Add(key,value)
}
func (this *HashTable)hash(key interface{}) int{
	return int(HashCode(key) & 0x7fffffff) %this.M
}
func  (this *HashTable)resize(newM int) {
	var newHashTable []*red_black_tree.RBTree
	for i:=0;i<newM;i++ {
		newHashTable = append(newHashTable,red_black_tree.NewRBTree())
	}
	oldM := this.M
	this.M = newM
	for i:=0;i<oldM ;i++  {
		m := this.hashtable[i]
		for _,key := range m.KeySet() {
			newHashTable[this.hash(key)].Add(key,m.Get(key))
		}
	}
	this.hashtable = newHashTable
}

func HashCode(source interface{}) uint64 {
	switch source.(type) {
	case int:
		source = strconv.Itoa(source.(int))
	case float64:
		source = strconv.FormatFloat(source.(float64), 'f', 6, 64)
	}

	h := fnv.New64a()
	h.Write([]byte(source.(string)))
	return h.Sum64()
}