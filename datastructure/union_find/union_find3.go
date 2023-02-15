package union_find

type UnionFind3 struct {
	parent []int
	sizes  []int // sizes[i]表示以i为根的集合中元素个数
}

func NewUnionFind3(size int) *UnionFind3 {
	parent := make([]int, size)
	sizes := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		sizes[i] = 1
	}
	return &UnionFind3{parent: parent, sizes: sizes}
}

func (this *UnionFind3) GetSize() int {
	return len(this.parent)
}

func (this *UnionFind3) find(p int) int {
	if p < 0 || p > len(this.parent) {
		panic("Invalid p is illegal")
	}
	// 不断去查询自己的父亲节点, 直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != this.parent[p] {
		p = this.parent[p]
	}
	return p
}

func (this *UnionFind3) IsConnected(p, q int) bool {
	return this.find(p) == this.find(q)
}

func (this *UnionFind3) UnionElements(p, q int) {
	pRoot := this.find(p)
	qRoot := this.find(q)
	if pRoot == qRoot {
		return
	}
	if this.sizes[pRoot] < this.sizes[qRoot] {
		this.parent[pRoot] = qRoot
		this.sizes[qRoot] += this.sizes[pRoot]
	} else {
		this.parent[qRoot] = pRoot
		this.sizes[pRoot] += this.sizes[qRoot]
	}
}
