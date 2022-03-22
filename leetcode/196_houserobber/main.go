package main

import "fmt"

/*
你是一个专业的小偷，计划偷窃沿街的房屋。

===> 每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统。
===> 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
*/
// https://leetcode-cn.com/problems/house-robber/

// dp table
func Rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// dp[i] 表示从第 i 间房开始盗窃, 能获取到的最高金额
	// base case: dp[n+1] = dp[n] = 0
	dp := make([]int, n+2)
	for i := n - 1; i >= 0; i-- {
		// dp[i] = Max(选择一, 选择二)
		// 选择一: 不抢当前房间, 所获金额 与后一个房间一致（下一个房间的选择）
		// 选择二: 抢劫当前房间, 所有金额 = 当前房间金额 + 往后数第二个房间所获金额
		dp[i] = Max(dp[i+1], nums[i]+dp[i+2])
	}

	return dp[0]
}

// dp table 空间复杂度优化
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp_i_1, dp_i_2 := 0, 0
	dp_i := 0
	for i := n - 1; i >= 0; i-- {
		dp_i = Max(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println("nums:[1,2,3,1] => 4", rob([]int{1, 2, 3, 1}))
	fmt.Println("nums:[2,7,9,3,1] => 12", rob([]int{2, 7, 9, 3, 1}))
	fmt.Println("nums:[1] => 1", rob([]int{1}))
}
