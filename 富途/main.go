package main

// 实现梯度收费

type FeeTier struct {
	MinOrders int
	MaxOrders int
	FeePer    int
}

func CalculateFee(orderCount int) int {
	tiers := []FeeTier{
		{1, 10, 20},
		{11, 100, 10},
		{101, 500, 5},
		{501, -1, 2}, // -1 表示无上限
	}

	totalFee := 0
	remaining := orderCount

	for _, tier := range tiers {
		if remaining <= 0 {
			break
		}

		var ordersInTier int
		if tier.MaxOrders == -1 {
			ordersInTier = remaining
		} else {
			maxPossible := tier.MaxOrders - tier.MinOrders + 1
			ordersInTier = min(remaining, maxPossible)
		}

		totalFee += ordersInTier * tier.FeePer
		remaining -= ordersInTier
	}

	return totalFee
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
