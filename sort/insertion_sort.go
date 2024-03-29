package sort

func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return
}

func InsertionSortAdvance(arr []int) {
	for i := 0; i < len(arr); i++ {
		t := arr[i]
		j := i
		for ; j > 0 && t < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = t
	}
	return
}
