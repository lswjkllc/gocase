package main

import (
	"fmt"
	"math"
)

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

===> 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
===> 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
===> k = 2 表示最多可以进行 2 次交易
*/
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

// 框架版本
func MaxProfit(prices []int) int {
	pLen := len(prices)
	if pLen < 2 {
		return 0
	}
	// 2次交易
	const maxK int = 2
	// dp 数组记录状态信息
	dp := make([][maxK + 1][2]int, pLen)
	for i := 0; i < pLen; i++ {
		for k := maxK; k >= 1; k-- {
			if i-1 < 0 {
				// 处理 base case
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = Max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = Max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[pLen-1][maxK][0]
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 空间复杂度优化
func maxProfit(prices []int) int {
	pLen := len(prices)
	if pLen < 2 {
		return 0
	}
	// 四个状态信息
	dp_i10, dp_i11 := 0, math.MinInt32
	dp_i20, dp_i21 := 0, math.MinInt32
	for i := 0; i < pLen; i++ {
		// 1 次交易
		dp_i11 = Max(dp_i11, -prices[i])
		dp_i10 = Max(dp_i10, dp_i11+prices[i])
		// 2 次交易
		dp_i21 = Max(dp_i21, dp_i10-prices[i])
		dp_i20 = Max(dp_i20, dp_i21+prices[i])
	}

	return dp_i20
}

func main() {
	fmt.Println("prices:[3, 3, 5, 0, 0, 3, 1, 4] k:2 => 6", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println("prices:[1, 2, 3, 4, 5], k:2 => 4", maxProfit([]int{1, 2, 3, 4, 5}))
}
