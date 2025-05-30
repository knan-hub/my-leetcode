package main

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择某一天买入这只股票，并选择在未来的某一个不同的日子卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

func maxProfit(prices []int) int {
	// 方法一
	// maxProfit := 0
	// for i := 0; i < len(prices); i++ {
	// 	profit := 0
	// 	for j := i + 1; j < len(prices); j++ {
	// 		if prices[j] > prices[i] {
	// 			profit = max(profit, prices[j]-prices[i])
	// 		}
	// 	}
	// 	maxProfit = max(maxProfit, profit)
	// }
	// return maxProfit

	// 方法二
	// minPrice := math.Inf(1)
	// maxProfit := 0.0
	// for _, price := range prices {
	// 	p := float64(price)
	// 	if p < minPrice {
	// 		minPrice = p
	// 	} else {
	// 		maxProfit = math.Max(maxProfit, p-minPrice)
	// 	}
	// }
	// return int(maxProfit)

	// 方法三
	// dp := make([][2]int, len(prices))
	// dp[0][0] = 0
	// dp[0][1] = -prices[0]
	// for i := 1; i < len(prices); i++ {
	// 	dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	// 	dp[i][1] = max(dp[i-1][1], -prices[i])
	// }
	// return dp[len(prices)-1][0]

	// 方法四
	cash := 0
	hold := -prices[0]
	for i := 1; i < len(prices); i++ {
		cash = max(cash, hold+prices[i])
		hold = max(hold, -prices[i])
	}
	return cash
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
