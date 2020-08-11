package test

import (
	"data-structures/algorithm"
	"testing"
)

func TestSort(t *testing.T) {
	strarr := algorithm.StringArray{"x", "A", "B", "E", "H", "G"}
	intArr := algorithm.IntArray{6, 3, 4, 8, 1}
	t.Log(algorithm.SelectionSort(strarr))
	//t.Log(algorithm.SelectionSort(intArr))
	t.Log(algorithm.InsertionSort(intArr))
}
