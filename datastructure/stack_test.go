
package datastructure

import (
	"testing"
)

func TestStack(t *testing.T) {
	l := NewLinkedListStack()
	for i := 0; i < 5; i++ {
		l.Push(i)
	}

	t.Log(l.Print())
	t.Log(l.Pop())
	t.Log(l.Print())
	t.Log(l.Peek())
}
