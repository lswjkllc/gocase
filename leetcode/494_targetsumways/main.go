package main

import "fmt"

/*
给你一个整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加 '+' 或 '-', 然后串联起所有整数，可以构造一个 表达式:
	例如, nums = [2, 1], 可以在 2 之前添加 '+', 在 1 之前添加 '-', 然后串联起来得到表达式 "+2-1" 。

返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
*/
// https://leetcode-cn.com/problems/target-sum/

var res = 0

func findTargetSumWays(nums []int, target int) int {
	res = 0
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 回溯算法
	backtrack(nums, 0, target)
	return res
}

// 回溯算法
func backtrack(nums []int, i int, rest int) {
	if i == len(nums) {
		if rest == 0 {
			res += 1
		}
		return
	}

	// 加 - 号
	rest += nums[i]
	backtrack(nums, i+1, rest)
	rest -= nums[i]

	// 加 + 号
	rest -= nums[i]
	backtrack(nums, i+1, rest)
	rest += nums[i]
}

// 消除重叠子问题

// 动态规划

func main() {
	fmt.Println("nums:[1,1,1,1,1] target:3 => 5", findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println("nums:[1] target:1 => 1", findTargetSumWays([]int{1}, 1))
}
