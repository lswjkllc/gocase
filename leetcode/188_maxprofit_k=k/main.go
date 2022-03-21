package main

import (
	"fmt"
	"math"
)

/*
给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。

===> 设计一个算法来计算你所能获取的最大利润。
===> 你最多可以完成 k 笔交易。
===> 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）
===> k = k 表示最多完成指定 k 次交易
*/
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/

// 框架版本
func maxProfit(k int, prices []int) int {
	pLen := len(prices)
	if pLen < 2 {
		return 0
	}
	// 一次交易由买入和卖出构成, 至少需要两天。
	// 所以说有效的限制k应该不超过 pLen/2。如果超过, 就没有约束作用了, 相当于k = +infinity。
	// 直接使用 第 122 题的方法解决
	if k > pLen/2 {
		return maxProfitKinf(prices)
	}

	// 如果 k 不超过 pLen/2, 可通过框架求解
	dp := make([][][2]int, pLen)
	// 分配 二维 空间, 同时处理 k = 0 的 base case
	for i := 0; i < pLen; i++ {
		dp[i] = make([][2]int, k+1)
	}

	for i := 0; i < pLen; i++ {
		for j := k; j >= 1; j-- {
			if i-1 < 0 {
				// 处理 base case
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = Max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[pLen-1][k][0]
}

func maxProfitKinf(prices []int) int {
	pLen := len(prices)
	if pLen < 2 {
		return 0
	}
	dp_i_0, dp_i_1 := 0, math.MinInt32
	for i := 0; i < pLen; i++ {
		// 因为 k != 1, 需要保存上一轮的 dp_i_0（sell结果）
		temp := dp_i_0
		// sell
		dp_i_0 = Max(dp_i_0, dp_i_1+prices[i])
		// buy
		dp_i_1 = Max(dp_i_1, temp-prices[i])
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
	fmt.Println("prices:[2, 4, 1] k:2 => 2", maxProfit(2, []int{2, 4, 1}))
	fmt.Println("prices:[3, 2, 6, 5, 0, 3] k:2 => 7", maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
	fmt.Println("prices:[1,2,4,2,5,7,2,4,9,0] k:4 => 15", maxProfit(4, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
}
