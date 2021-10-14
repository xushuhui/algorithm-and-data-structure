package sort

func BubbleSort(nums []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := 0; j <l-1-i ; j++ {
			if nums[j] > nums[j+1]{
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
	return 
}

func BubbleSort2(nums []int)  {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		var isSort = true
		for j := 0; j <l-1-i ; j++ {
			if nums[j] > nums[j+1]{
				nums[j],nums[j+1] = nums[j+1],nums[j]
				isSort = false
			}
		}
		if isSort {
			break
		}
	}
	return 
}
func BubbleSort3(nums []int)  {
	l := len(nums)
	lastIndex := 0
	sortBorder := l -1

	for i := 0; i < l-1; i++ {
		var isSort = true
		for j := 0; j <sortBorder ; j++ {
			if nums[j] > nums[j+1]{
				nums[j],nums[j+1] = nums[j+1],nums[j]
				isSort = false
				lastIndex = j
			}

		}
		sortBorder = lastIndex
		if isSort {
			break
		}
	}
	return 
}