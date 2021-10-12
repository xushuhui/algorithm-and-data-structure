package sort

import (
	"data-structures/utils"
	"testing"
)

func TestSort(t *testing.T) {
	t.Log(BubbleSort([]int{2,3,1,5,4,6}))
	//n := 10000000
	//s1 := utils.GenerateRandomArray(n, 0, n)
	//s2 := utils.CopyArray(s1, n)
	//s3 := utils.CopyArray(s1, n)

	//utils.TimeSpent("SelectionSort", algorithm.SelectionSort, s1)
	//utils.TimeSpent("InsertionSort", algorithm.InsertionSort, s2)
	//utils.TimeSpent("InsertionSortAdvance", algorithm.InsertionSortAdvance, s3)

	//utils.TimeSpent("SelectionSort", algorithm.SelectionSort, s4)

	// for _, v := range []int{100000, 1000000, 10000000} {
	// 	s := utils.GenerateOrderedArray(v)
	// 	sort(s, v)
	// }
	// for _, v := range []int{100000, 1000000, 10000000} {
	// 	s := utils.GenerateTestOrderedArray(v)
	// 	sort(s, v)
	// }
}
func sort(s []int, n int) {

	s5 := utils.CopyArray(s, n)
	s6 := utils.CopyArray(s, n)

	//utils.TimeSpent("SelectionSort", algorithm.SelectionSort, s4)
	utils.TimeSpent("InsertionSort", InsertionSort, s5)
	utils.TimeSpent("InsertionSortAdvance", InsertionSortAdvance, s6)
}
