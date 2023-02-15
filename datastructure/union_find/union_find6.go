package union_find

type UnionFind6 struct {
	parent []int
	rank   []int // rank[i]表示以i为根的集合所表示的树的层数
}

func NewUnionFind6(size int) *UnionFind6 {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind6{parent: parent, rank: rank}
}

func (this *UnionFind6) GetSize() int {
	return len(this.parent)
}

func (this *UnionFind6) find(p int) int {
	if p < 0 || p > len(this.parent) {
		panic("Invalid p is illegal")
	}
	// 不断去查询自己的父亲节点, 直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != this.parent[p] {
		this.parent[p] = this.parent[this.parent[p]]
		p = this.parent[p]
	}
	if p != this.parent[p] {
		this.parent[p] = this.find(this.parent[p])
	}
	return p
}

func (this *UnionFind6) IsConnected(p, q int) bool {
	return this.find(p) == this.find(q)
}

func (this *UnionFind6) UnionElements(p, q int) {
	pRoot := this.find(p)
	qRoot := this.find(q)
	if pRoot == qRoot {
		return
	}
	// 根据两个元素所在树的rank不同判断合并方向
	// 将rank低的集合合并到rank高的集合上
	if this.rank[pRoot] < this.rank[qRoot] {
		this.parent[pRoot] = qRoot
	} else if this.rank[qRoot] < this.rank[pRoot] {
		this.parent[qRoot] = pRoot
	} else {
		this.parent[pRoot] = qRoot
		this.rank[qRoot] += 1
	}
}
