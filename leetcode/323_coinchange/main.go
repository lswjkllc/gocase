package main

import "fmt"

func coinChange(coins []int, amount int) int {
	// base case
	cLen := len(coins)
	if cLen < 1 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	// dp[i] 表示 目标金额 amuout=i 时, 需要的 coin 数量
	// dp[0] = 0
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = Min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("coins:[1, 2, 5] amount:11 => 3", coinChange([]int{1, 2, 5}, 11))
	fmt.Println("coins:[2] amount:3 => -1", coinChange([]int{2}, 3))
	fmt.Println("coins:[1] amount:0 => 0", coinChange([]int{1}, 0))
}
