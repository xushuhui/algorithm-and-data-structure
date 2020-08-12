package test

import (
	"data-structures/algorithm"
	"data-structures/utils"
	"testing"
)

func TestSort(t *testing.T) {
	n := 100000
	s1 := utils.GenerateRandomArray(n, 0, n)
	//s1 := []interface{}{6,4,1,3,4,2,8,7}
	s2 := utils.CopyArray(s1, n)
	s3 := utils.CopyArray(s1, n)

	utils.TimeSpent("SelectionSort", algorithm.SelectionSort, s1)
	utils.TimeSpent("InsertionSort", algorithm.InsertionSort, s2)
	utils.TimeSpent("InsertionSortAdvance", algorithm.InsertionSortAdvance, s3)

}
