package algorithm

import "sort"

func InsertionSort(arr sort.Interface) sort.Interface {
	for i := 0; i < arr.Len(); i++ {
		for j := i; j >= 1; j-- {
			if arr.Less(j, j-1) {
				arr.Swap(j, j-1)
			}
		}
	}
	return arr
}
