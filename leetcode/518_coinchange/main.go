package main

import "fmt"

/*
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

假设每一种面额的硬币有无限个。

题目数据保证结果符合 32 位带符号整数。
*/
// https://leetcode-cn.com/problems/coin-change-2/

func Change(amount int, coins []int) int {
	N := len(coins)
	W := amount

	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i <= N; i++ {
		for j := 0; j <= W; j++ {
			if i == 0 {
				dp[i][j] = 0
				continue
			}
			if j == 0 {
				dp[i][j] = 1
				continue
			}

			dp[i][j] = dp[i-1][j]
			if j-coins[i-1] >= 0 {
				dp[i][j] += dp[i][j-coins[i-1]]
			}
			// if j-coins[i-1] >= 0 {
			// 	dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			// } else {
			// 	dp[i][j] = dp[i-1][j]
			// }
		}
	}

	return dp[N][W]
}

// 空间复杂度优化
func change(amount int, coins []int) int {
	N := len(coins)
	W := amount

	dp := make([]int, W+1)
	dp[0] = 1

	for i := 0; i < N; i++ {
		for j := 1; j <= W; j++ {
			if j-coins[i] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}

	return dp[W]
}

func main() {
	fmt.Println("coins:[1, 2, 5] amount:5 => 4", change(5, []int{1, 2, 5}))
	fmt.Println("coins:[2] amount:3 => 0", change(3, []int{2}))
}
