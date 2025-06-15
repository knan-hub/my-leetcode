package main

// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

func majorityElement(nums []int) int {
	// 方法一
	// memo := make(map[int]int, len(nums))
	// for _, num := range nums {
	// 	memo[num]++
	// 	if memo[num] > len(nums)/2 {
	// 		return num
	// 	}
	// }
	// return -1

	// 方法二
	candidate := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
			if count == 0 {
				candidate = nums[i]
				count = 1
			}
		}
	}
	return candidate
}
