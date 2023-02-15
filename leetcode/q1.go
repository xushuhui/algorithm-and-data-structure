package leetcode

/*
第一题：寻找缺失的数
题目大意：
	一个长度为 n-1 的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围 0～n-1 之内。
	在范围 0～n-1 内的 n 个数字中有且只有一个数字不在该数组中，请找出这个数字。
示例 1:
	输入: [0,1,3]
	输出: 2
示例 2:
	输入: [0,1,2,3,4,5,6,7,9]
	输出: 8
限制：
	1 <= 数组长度 <= 10000
思考：能否在你确定算法之后，优化一下*/
//时间复杂度O(n) 空间复杂度O(n)
func q1(nums []int) int {
	l := len(nums)
	if l < 1 || l > 10000 {
		panic(" index out of slice ")
	}
	// 另一个数组
	b := make([]int, 0, l)
	for i := 0; i < l; i++ {
		b = append(b, i)
	}
	var target int
	for i, v := range b {
		if v != nums[i] {
			target = v
		}
	}

	return target
}

// 时间复杂度O(n) 空间复杂度O(1)
func q1_1(nums []int) int {
	var target int

	for i, v := range nums {
		if i != v {
			target = i
		}
	}
	return target
}

func q1_2(nums []int) int {
	l := len(nums)
	if l < 1 || l > 10000 {
		panic(" index out of slice ")
	}
	left := 0
	right := l - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
