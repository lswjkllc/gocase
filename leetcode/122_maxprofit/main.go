package main

import "fmt"

/*
给定一个数组 prices ，其中 prices[i] 表示股票第 i 天的价格。

===> 在每一天，你可能会决定购买和/或出售股票。
===> 你在任何时候 最多 只能持有 一股 股票。
===> 你也可以购买它，然后在 同一天 出售。

返回 你能获得的 最大 利润 。
*/
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

// 该方法在 prices 很大的时候，会超过 leetcode 给的时间限制
func MaxProfit(prices []int) int {
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
func maxProfit(prices []int) int {
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
