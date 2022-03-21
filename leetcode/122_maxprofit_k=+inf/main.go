package main

import (
	"fmt"
	"math"
)

/*
给定一个数组 prices ，其中 prices[i] 表示股票第 i 天的价格。

===> 在每一天，你可能会决定购买和/或出售股票。
===> 你在任何时候 最多 只能持有 一股 股票。
===> 你也可以购买它，然后在 同一天 出售。（当天卖出没有利润，所以这个条件没有意义）
===> k = +infinity 表示不限制交易次数

返回 你能获得的 最大 利润 。
*/
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

// 该方法在 prices 很大的时候，会超过 leetcode 给的时间限制
func MaxProfit1(prices []int) int {
	pLen := len(prices)
	memo := map[int]int{}
	for i := 0; i < pLen; i++ {
		memo[i] = -1
	}
	var dp func(int) int
	dp = func(start int) int {
		if start >= pLen {
			return 0
		}
		if v, ok := memo[start]; ok && v != -1 {
			return v
		}
		res := 0
		curMin := prices[start]
		for sell := start + 1; sell < pLen; sell++ {
			curMin = Min(curMin, prices[sell])
			res = Max(res, dp(sell+1)+prices[sell]-curMin)
		}

		return res
	}
	return dp(0)
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 最优解：贪心算法
func MaxProfit2(prices []int) int {
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}

// 框架版本
func MaxProfit3(prices []int) int {
	pLen := len(prices)
	if pLen < 2 {
		return 0
	}
	/*
		如果k为正无穷，那么就可以认为k和k - 1是一样的。框架：
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
		            	= max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])
		我们发现数组中的 k 已经不会改变了，也就是说不需要记录 k 这个状态了：
			dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])

	*/
	dp := make([][2]int, pLen)
	for i := 0; i < pLen; i++ {
		if i-1 < 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
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

func main() {
	fmt.Println("prices:[7, 1, 5, 3, 6, 4] k:+inf => 7", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("prices:[1, 2, 3, 4, 5] k:+inf => 4", maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println("prices:[7, 6, 4, 3, 1] k:+inf => 0", maxProfit([]int{7, 6, 4, 3, 1}))
}
