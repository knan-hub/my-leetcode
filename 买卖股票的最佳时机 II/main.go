package main

// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候最多只能持有一股股票。你也可以先购买，然后在同一天出售。
// 返回你能获得的最大 利润 。

func maxProfit(prices []int) int {
	// 方法一：贪心算法
	// profit := 0
	// for i := 1; i < len(prices); i++ {
	// 	if prices[i] > prices[i-1] {
	// 		profit += prices[i] - prices[i-1]
	// 	}
	// }
	// return profit

	// 方法二：动态规划 二维数组
	// dp := make([][2]int, len(prices))
	// dp[0][0] = 0
	// dp[0][1] = -prices[0]
	// for i := 1; i < len(prices); i++ {
	//     // dp[i][0] 表示第 i 天结束后不持股的最大利润
	//     // dp[i][1] 表示第 i 天结束后持股的最大利润
	// 	dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	// 	dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	// }
	// return dp[len(prices)-1][0]

	// 方法三：动态规划 滚动变量
	cash := 0          // 手里没股票
	hold := -prices[0] // 手里有股票
	for i := 1; i < len(prices); i++ {
		cash = max(cash, hold+prices[i])
		hold = max(hold, cash-prices[i])
	}
	return cash
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
