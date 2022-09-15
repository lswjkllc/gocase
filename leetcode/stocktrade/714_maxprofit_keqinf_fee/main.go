package main

import (
	"fmt"
	"math"
)

/*
给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；
整数 fee 代表了交易股票的手续费用。

===> 你可以无限次地完成交易，但是你每笔交易都需要付手续费。
===> 如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
===> 返回获得利润的最大值。
===> 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
===> k = +infinity 表示不限制交易次数

*/
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

// 框架版本
func MaxProfit(prices []int, fee int) int {
	pLen := len(prices)
	if pLen < 2 {
		return 0
	}
	/*
		每次交易要支付手续费，即把手续费从利润中减去。把这个特点融入 第 122 题 的状态转移方程即可:
			dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)
		解释：相当于买入股票的价格升高了。
		在第一个式子里减也是一样的，相当于卖出股票的价格减小了。
		如果直接把fee放在第一个式子里减，可能会有测试用例无法通过，错误原因是整型溢出而不是思路问题。
		一种解决方案是把代码中的int类型都改成long类型，避免int的整型溢出。
	*/
	dp := make([][2]int, pLen)
	for i := 0; i < pLen; i++ {
		if i-1 < 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i] - fee
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	return dp[pLen-1][0]
}

// 空间复杂度优化版本
func maxProfit(prices []int, fee int) int {
	pLen := len(prices)
	if pLen < 2 {
		return 0
	}
	dp_i_0, dp_i_1 := 0, math.MinInt32
	for i := 0; i < pLen; i++ {
		temp := dp_i_0
		dp_i_0 = Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = Max(dp_i_1, temp-prices[i]-fee)
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
	fmt.Println("prices:[1, 3, 2, 8, 4, 9] k:+inf fee:2 => 8", maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println("prices:[1, 3, 7, 5, 10, 3] k:+inf fee:3 => 6", maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}
