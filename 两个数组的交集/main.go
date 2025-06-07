package main

// 给定两个数组 nums1 和 nums2 ，返回它们的交集 。输出结果中的每个元素一定是唯一的。我们可以不考虑输出结果的顺序。

func intersection(nums1 []int, nums2 []int) []int {
	memo := map[int]struct{}{}
	for _, num := range nums1 {
		memo[num] = struct{}{}
	}
	resultMemo := map[int]struct{}{}
	for _, num := range nums2 {
		if _, ok := memo[num]; ok {
			resultMemo[num] = struct{}{}
		}
	}
	result := []int{}
	for num := range resultMemo {
		result = append(result, num)
	}
	return result
}
