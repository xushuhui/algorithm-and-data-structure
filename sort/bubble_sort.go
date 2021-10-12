package sort

func BubbleSort(nums []int)[]int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := l-1; j >i ; j-- {
			if nums[j-1] > nums[j]{
				nums[j],nums[j-1] = nums[j-1],nums[j]
			}
		}
	}
	return nums
}
