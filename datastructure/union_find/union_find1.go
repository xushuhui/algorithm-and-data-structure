package union_find

type UnionFind1 struct {
	id []int
}

func NewUnionFind1(size int) *UnionFind1 {
	id := make([]int, size)
	for i := 0; i < size; i++ {
		id[i] = i
	}
	return &UnionFind1{id: id}
}

func (this *UnionFind1) GetSize() int {
	return len(this.id)
}

func (this *UnionFind1) find(p int) int {
	if p < 0 || p > len(this.id) {
		panic("Invalid p is illegal")
	}
	return this.id[p]
}

func (this *UnionFind1) IsConnected(p, q int) bool {
	return this.find(p) == this.find(q)
}

func (this *UnionFind1) UnionElements(p, q int) {
	pid := this.find(p)
	qid := this.find(q)
	if pid == qid {
		return
	}
	for i := 0; i < len(this.id); i++ {
		if this.id[i] == pid {
			this.id[i] = qid
		}
	}
}
