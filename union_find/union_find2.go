type UnionFind2 struct {
	parent  []int
}
func NewUnionFind2(size int)*UnionFind2 {
	parent  = make([]int, size)
	for i := 0; i < size; i++ {
		parent [i] = i
	}
	return &UnionFind2{parent : parent }
}
func (this *UnionFind2))GetSize()int{
	return len(this.parent )
}

func (this *UnionFind2)find(p int)int{
	if p <0 || p > len(this.parent ) {
		panic("Invalid is illegal")
	}
	// 不断去查询自己的父亲节点, 直到到达根节点
		// 根节点的特点: parent[p] == p
		for p != this.parent[p] {
			p = this.parent[p]
		}
	return p
}
func (this *UnionFind2)IsConnected(p,q int)bool{
	return this.find(p)==this.find(q)
}
func (this *UnionFind2)UnionElements(p,q int){
	pid := this.find(p)
	qid := this.find(q)
	if pid==qid{
		return
	}
	this.parent[pid] =qid
}