package sort

import (
	"data-structures/utils"
)

func InsertionSort(arr []interface{}) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && utils.Compare(arr[j], arr[j-1]) < 0; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return
}
func InsertionSortAdvance(arr []interface{}) {
	for i := 0; i < len(arr); i++ {
		t := arr[i]
		j := i
		for ; j > 0 && utils.Compare(t, arr[j-1]) < 0; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = t
	}
	return
}
