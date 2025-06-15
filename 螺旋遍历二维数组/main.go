package main

// 给定一个二维数组 array，请返回「螺旋遍历」该数组的结果。
// 螺旋遍历：从左上角开始，按照 向右、向下、向左、向上 的顺序 依次 提取元素，然后再进入内部一层重复相同的步骤，直到提取完所有元素。

func spiralArray(array [][]int) []int {
	if len(array) == 0 || len(array[0]) == 0 {
		return []int{}
	}
	top, bottom := 0, len(array)-1
	left, right := 0, len(array[0])-1
	res := []int{}
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, array[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			res = append(res, array[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, array[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, array[i][left])
			}
			left++
		}
	}
	return res
}
