package main

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target  的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
// 你可以按任意顺序返回答案。

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		if value, ok := mapping[num]; ok {
			return []int{i, value}
		}
		mapping[nums[i]] = i
	}
	return []int{}
}
