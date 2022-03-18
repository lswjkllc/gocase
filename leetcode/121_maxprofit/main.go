package main

import "fmt"

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

===> 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。

设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*/
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	res, pLen := 0, len(prices)
	if pLen == 0 {
		return res
	}
	curMin := prices[0]
	for sell := 1; sell < pLen; sell++ {
		curMin = min(curMin, prices[sell])
		res = max(res, prices[sell]-curMin)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
