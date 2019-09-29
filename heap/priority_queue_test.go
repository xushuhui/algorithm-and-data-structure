package heap

import "testing"

func TestPriorityQueue(t *testing.T) {
	q := PriorityQueueConstructor()
	for i:=0;i<=10 ;i++  {
		q.Enqueue(i)
	}
	t.Log(q.maxHeap.data)
}