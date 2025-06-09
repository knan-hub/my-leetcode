package main

// 给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。
// 返回满足此条件的任一数组作为答案。

func sortArrayByParity(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		for left < right && nums[left]%2 == 0 {
			left++
		}
		for left < right && nums[right]%2 != 0 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
