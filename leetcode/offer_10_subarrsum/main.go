package main

import "fmt"

/*
给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
*/
// https://leetcode-cn.com/problems/QTMn0o/

func SubarraySum(nums []int, k int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}

	preSum := make([]int, N+1)
	for i := 1; i <= N; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	ans := 0
	for i := 1; i <= N; i++ {
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] == k {
				ans += 1
			}
		}
	}

	return ans
}

// 时间复杂度优化
func subarraySum(nums []int, k int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}

	// 前缀和 map: 前缀和 -> 前缀和出现的次数
	preSum := make(map[int]int)
	// base case
	preSum[0] = 1

	ans, sum_i := 0, 0
	for i := 0; i < N; i++ {
		// 到目前为止的前缀和 nums[0..i]
		sum_i += nums[i]
		// 每一轮想找的前缀和 nums[0..j]
		sum_j := sum_i - k
		// 如果该前缀和存在, 结果+1
		if v, ok := preSum[sum_j]; ok {
			ans += v
		}
		// 更新前缀和 map
		preSum[sum_i] = preSum[sum_i] + 1
	}

	return ans
}

func main() {
	fmt.Println("nums:[1,1,1] k:2 => 2", subarraySum([]int{1, 1, 1}, 2))
	fmt.Println("nums:[1,2,3] k:3 => 2", subarraySum([]int{1, 2, 3}, 3))
}
