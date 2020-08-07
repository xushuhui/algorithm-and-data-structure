package test

import (
	"data-structures/algorithm"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{2, 3, 1, 6, 4, 8}
	t.Log(algorithm.BubbleSort(arr))
}
