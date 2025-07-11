package main

// 给定一个未经排序的整数数组，找到最长且连续递增的子序列，并返回该序列的长度。
// 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSize := 1
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] > nums[j-1] {
			size := j - i + 1
			maxSize = max(maxSize, size)
		} else {
			i = j
		}
	}
	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
