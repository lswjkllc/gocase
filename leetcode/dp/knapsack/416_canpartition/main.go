package main

import "fmt"

/*
给你一个 只包含正整数 的 非空 数组 nums。
请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
*/
// https://leetcode-cn.com/problems/partition-equal-subset-sum/

func CanPartition(nums []int) bool {
	// 直接根据集合长度判断
	N := len(nums)
	if N == 1 {
		return false
	}
	// 计算 和
	sum := 0
	for i := 0; i < N; i++ {
		sum += nums[i]
	}
	// 和 不能平分, 直接返回 false
	if sum%2 != 0 {
		return false
	}

	W := sum / 2

	dp := make([][]bool, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]bool, W+1)
	}
	// dp[i][0] = true, dp[0][w] = false
	for i := 1; i <= N; i++ {
		for w := 0; w <= W; w++ {
			if w == 0 {
				dp[i][w] = true
				continue
			}
			if w < nums[i-1] {
				dp[i][w] = dp[i-1][w]
				continue
			}
			dp[i][w] = dp[i-1][w] || dp[i-1][w-nums[i-1]]
		}
	}

	return dp[N][W]
}

// 空间复杂度优化
func canPartition(nums []int) bool {
	// 直接根据集合长度判断
	N := len(nums)
	if N == 1 {
		return false
	}
	// 计算 和
	sum := 0
	for i := 0; i < N; i++ {
		sum += nums[i]
	}
	// 和 不能平分, 直接返回 false
	if sum%2 != 0 {
		return false
	}

	W := sum / 2

	dp := make([]bool, W+1)
	dp[0] = true
	for i := 1; i <= N; i++ {
		for w := W; w >= 0; w-- {
			if w >= nums[i-1] {
				dp[w] = dp[w] || dp[w-nums[i-1]]
			}
		}
	}

	return dp[W]
}

func main() {
	fmt.Println("nums:[1,5,11,5] => true", canPartition([]int{1, 5, 11, 5}))
	fmt.Println("nums:[1,2,3,5] => false", canPartition([]int{1, 2, 3, 5}))
}
