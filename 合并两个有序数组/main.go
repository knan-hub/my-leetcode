package main

// 给你两个按非递减顺序排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
// 请你合并 nums2 到 nums1 中，使合并后的数组同样按非递减顺序排列。
// 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

func merge(nums1 []int, m int, nums2 []int, n int) {
	indexM := m - 1
	indexN := n - 1
	index := m + n - 1
	for indexM >= 0 && indexN >= 0 {
		if nums1[indexM] < nums2[indexN] {
			nums1[index] = nums2[indexN]
			indexN--
		} else {
			nums1[index] = nums1[indexM]
			indexM--
		}
		index--
	}
	for indexN >= 0 {
		nums1[index] = nums2[indexN]
		indexN--
		index--
	}
}
