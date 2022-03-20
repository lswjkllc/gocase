package main

import (
	"fmt"
	"math"
)

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

===> 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。
===> k = 1 表示只能进行 1 次交易

设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*/
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

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

func MaxProfit1(prices []int) int {
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

func MaxProfit2(prices []int) int {
	pLen := len(prices)
	if pLen <= 1 {
		return 0
	}
	// dp[n][k][2]
	// k = 1 => dp[n][2]
	// k - 1 = 0 => dp[i-1][0]
	/*
		k = 0 的 base case，所以 dp[i-1][0][0] = 0。
		现在发现 k 都是 1，不会改变，即 k 对状态转移已经没有影响了。
		可以进行进一步化简去掉所有 k
	*/
	dp := make([][2]int, pLen)
	for i := 0; i < pLen; i++ {
		if i-1 < 0 {
			// 根据状态转移方程可得：
			//   dp[i][0]
			// = max(dp[-1][0], dp[-1][1] + prices[i])
			// = max(0, -infinity + prices[i]) = 0
			dp[i][0] = 0
			// 根据状态转移方程可得：
			//   dp[i][1]
			// = max(dp[-1][1], dp[-1][0] - prices[i])
			// = max(-infinity, 0 - prices[i])
			// = -prices[i]
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[pLen-1][0]
}

/*
注意一下状态转移方程，新状态只和相邻的一个状态有关。
其实不用整个dp数组，只需要一个变量储存相邻的那个状态就足够了，这样可以把空间复杂度降到 O(1)
*/
func MaxProfit(prices []int) int {
	pLen := len(prices)
	if pLen < 1 {
		return 0
	}
	dp_i_0, dp_i_1 := 0, math.MinInt32
	for i := 0; i < pLen; i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0
}

func main() {
	fmt.Println("prices: [7, 1, 5, 3, 6, 4] k: 1 => 5", MaxProfit([]int{7, 1, 5, 3, 6, 4}) == 5)
	fmt.Println("prices: [7, 6, 4, 3, 1] k: 1 => 0", MaxProfit([]int{7, 6, 4, 3, 1}) == 0)
}
