package array

import (
	"testing"
)

func TestArray(t *testing.T) {
	a := NewArray(8)
	for i := 0; i < 5; i++ {
		a.Add(i, i)
	}

	t.Log(a)
	a.AddV2(1, 2)
	t.Log(a)
}
