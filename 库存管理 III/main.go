package main

// 仓库管理员以数组 stock 形式记录商品库存表，其中 stock[i] 表示对应商品库存余量。请返回库存余量最少的 cnt 个商品余量，返回顺序不限。

func inventoryManagement(stock []int, cnt int) []int {
	sort(stock, 0, len(stock)-1, cnt)
	return stock[:cnt]
}

func sort(stock []int, left, right, cnt int) {
	if left >= right {
		return
	}
	index := quickSort(stock, left, right)
	if index == cnt {
		return
	} else if index > cnt {
		sort(stock, left, index-1, cnt)
	} else {
		sort(stock, index+1, right, cnt)
	}
}

func quickSort(stock []int, left, right int) int {
	newLeft := left
	newRight := right
	temp := stock[newLeft]
	for newLeft < newRight {
		for newLeft < newRight && stock[newRight] >= temp {
			newRight--
		}
		stock[newLeft] = stock[newRight]
		for newLeft < newRight && stock[newLeft] <= temp {
			newLeft++
		}
		stock[newRight] = stock[newLeft]
	}
	stock[newLeft] = temp
	return newLeft
}
