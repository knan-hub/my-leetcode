package main

// 展览馆展出来自 13 个朝代的文物，每排展柜展出 5 个文物。某排文物的摆放情况记录于数组 places，其中 places[i] 表示处于第 i 位文物的所属朝代编号。其中，编号为 0 的朝代表示未知朝代。请判断并返回这排文物的所属朝代编号是否能够视为连续的五个朝代（如遇未知朝代可算作连续情况）。

func checkDynasty(places []int) bool {
	// 方法一
	// sort.Ints(places)
	// zeroCount := 0
	// gaps := 0
	// for i := 0; i < len(places); i++ {
	// 	if places[i] == 0 {
	// 		zeroCount++
	// 	}
	// }
	// for i := zeroCount; i < len(places)-1; i++ {
	// 	if places[i] == places[i+1] {
	// 		return false
	// 	}
	// 	gaps += places[i+1] - places[i] - 1
	// }
	// return gaps <= zeroCount

	// 方法二
	memo := map[int]struct{}{}
	minNum, maxNum := 14, 0
	for _, v := range places {
		if v == 0 {
			continue
		}
		if _, ok := memo[v]; ok {
			return false
		}
		minNum = min(minNum, v)
		maxNum = max(maxNum, v)
		memo[v] = struct{}{}
	}
	return maxNum-minNum < 5
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
