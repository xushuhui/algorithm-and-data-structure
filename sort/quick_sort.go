package sort

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := Partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
	// index := Partition(arr,left,right)
	// quickSort()
}

func Partition(nums []int, left, right int) int {
	pivot := nums[left]

	for left != right {
		for left < right && nums[right] > pivot {
			right--
		}
		for left < right && nums[left] < pivot {
			left++
		}
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}

	}
	nums[left] = pivot
	return left
}
