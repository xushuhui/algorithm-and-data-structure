package sort



func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i; j < n; j++ {

			if arr[j]<  arr[minIndex]{
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	return
}
