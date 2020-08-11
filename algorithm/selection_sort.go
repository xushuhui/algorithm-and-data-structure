package algorithm

import "sort"

func SelectionSort(arr sort.Interface) sort.Interface {
	n := arr.Len()
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i; j < n; j++ {
			if arr.Less(j, minIndex) {
				minIndex = j
			}
		}
		arr.Swap(i, minIndex)
	}
	return arr
}
