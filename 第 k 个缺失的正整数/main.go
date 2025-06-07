package main

// 给你一个严格升序排列的正整数数组 arr 和一个整数 k 。
// 请你找到这个数组里第 k 个缺失的正整数。

func findKthPositive(arr []int, k int) int {
	// 方法一
	// if k <= arr[0]-1 {
	// 	return k
	// }
	// k -= arr[0] - 1
	// for i := 1; i < len(arr); i++ {
	// 	missing := arr[i] - arr[i-1] - 1
	// 	if k <= missing {
	// 		return arr[i-1] + k
	// 	}
	// 	k -= missing
	// }
	// return arr[len(arr)-1] + k

	// 方法二
	// 对于 arr[i]，前面应该有 arr[i] - 1 个正整数，但实际上只有 i 个在数组中。所以缺失的数量为：missing := arr[i] - (i + 1)
	left, right := 0, len(arr)-1
	for left <= right {
		middle := (left + right) >> 1
		missing := arr[middle] - (middle + 1)
		if missing < k {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left + k
}
