package main

import (
	"fmt"
	"math"
)

/*
给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​

===> 设计一个算法计算出最大利润。
===> 在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
===> 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
===> 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
===> k = +infinity 表示不限制交易次数

*/
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

// 框架版本
func MaxProfit(prices []int) int {
	pLen := len(prices)
	if pLen < 2 {
		return 0
	}
	/*
		每次sell之后要等一天才能继续交易。把这个特点融入 第 122 题 的状态转移方程即可:
			dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
		解释：第 i 天选择 buy 的时候，要从 i-2 的状态转移，而不是 i-1 。
	*/
	dp := make([][2]int, pLen)
	for i := 0; i < pLen; i++ {
		if i-1 < 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		if i-2 < 0 {
			dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = Max(dp[i-1][1], -prices[i])
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[pLen-1][0]
}

// 空间复杂度优化版本
func maxProfit(prices []int) int {
	pLen := len(prices)
	if pLen < 2 {
		return 0
	}
	dp_i_0, dp_i_1 := 0, math.MinInt32
	dp_pre_0 := 0 // 代表: dp[i-2][0]
	for i := 0; i < pLen; i++ {
		temp := dp_i_0
		dp_i_0 = Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = Max(dp_i_1, dp_pre_0-prices[i])
		dp_pre_0 = temp
	}
	return dp_i_0
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println("prices:[1, 2, 3, 0, 2] k:+inf cd:1 => 3", maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println("prices:[1] k:+inf cd:1 => 0", maxProfit([]int{1}))
}
