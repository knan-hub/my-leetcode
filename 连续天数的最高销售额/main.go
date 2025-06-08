package main

// 某公司每日销售额记于整数数组 sales，请返回所有连续一或多天销售额总和的最大值。
// 要求实现时间复杂度为 O(n) 的算法。

func maxSales(sales []int) int {
	// 方法一
	// if len(sales) == 0 {
	// 	return 0
	// }
	// maxSum := sales[0]
	// currentSum := sales[0]
	// // startIndex := 0
	// // endIndex := 0
	// // tempStart := 0
	// for i := 1; i < len(sales); i++ {
	// 	if currentSum < 0 {
	// 		currentSum = sales[i]
	// 		// tempStart = i
	// 	} else {
	// 		currentSum += sales[i]
	// 	}
	// 	if currentSum > maxSum {
	// 		maxSum = currentSum
	// 		// startIndex = tempStart
	// 		// endIndex = i
	// 	}
	// }
	// return maxSum

	// 方法二
	max := sales[0]
	for i := 1; i < len(sales); i++ {
		if sales[i]+sales[i-1] > sales[i] {
			sales[i] += sales[i-1]
		}
		if sales[i] > max {
			max = sales[i]
		}
	}
	return max
}
