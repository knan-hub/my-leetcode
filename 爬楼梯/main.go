package main

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

func climbStairs(n int) int {
	// 方法一
	// if n == 1 {
	// 	return 1
	// }
	// if n == 2 {
	// 	return 2
	// }
	// i := 1
	// j := 2
	// for k := 3; k <= n; k++ {
	// 	num := i + j
	// 	i = j
	// 	j = num
	// }
	// return j

	// 方法二
	// if n == 1 {
	// 	return 1
	// }
	// if n == 2 {
	// 	return 2
	// }
	// return climbStairs(n-1) + climbStairs(n-2)

	// 方法三
	memo := make(map[int]int)
	return helper(n, memo)
}

func helper(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if v, ok := memo[n]; ok {
		return v
	}
	memo[n] = helper(n-1, memo) + helper(n-2, memo)
	return memo[n]
}
