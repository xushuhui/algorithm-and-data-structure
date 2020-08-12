package algorithm

import (
	"data-structures/utils"
)

func SelectionSort(arr []interface{}) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i; j < n; j++ {

			if utils.Compare(arr[j], arr[minIndex]) < 0 {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	return
}
