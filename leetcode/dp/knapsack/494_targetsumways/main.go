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

func FindTargetSumWays(nums []int, target int) int {
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

/*
A 集合: 加 + 的集合
B 集合: 加 - 的集合

条件:
	sum(A) + sum(B) = sum(nums)
目标:
	sum(A) - sum(B) = target
推理:
	sum(A) = target + sum(B)
	sum(A) + sum(A) = target + sum(B) + sum(A)
	2 * sum(A) = target + sum(B) + sum(A)
	sum(A) = (target + sum(B) + sum(A)) / 2
	sum(A) = (target + sum(nums)) / 2
结论: 找到一个集合 A 的值等于 (目标值+集合值)/2

将这个问题转化为 背包问题, 可以使用动态规划方法解决
*/
// 动态规划（空间复杂度优化版本）
func findTargetSumWays(nums []int, target int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}
	if N == 1 {
		if nums[0]-target == 0 || nums[0]+target == 0 {
			return 1
		}
		return 0
	}
	sum := 0
	for i := 0; i < N; i++ {
		sum += nums[i]
	}
	if sum < target || (sum+target)%2 != 0 {
		return 0
	}

	W := (sum + target) / 2
	// 初始化 dp: dp[i] 表示 目标为 i 时, 具有多少种方法
	dp := make([]int, W+1)
	// 设置初值: 当目标为 0 时, 具有唯一一种方法（很重要）
	dp[0] = 1
	for i := 0; i < N; i++ {
		for w := W; w >= 0; w-- {
			v := nums[i]
			if w >= v {
				dp[w] = dp[w] + dp[w-v]
			}
		}
	}

	return dp[W]
}

func main() {
	fmt.Println("nums:[0,0,0,0,0,0,0,0,1] target:1 => 256", findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
	fmt.Println("nums:[1,1,1,1,1] target:3 => 5", findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println("nums:[1] target:1 => 1", findTargetSumWays([]int{1}, 1))
	fmt.Println("nums:[100] target:-200 => 0", findTargetSumWays([]int{100}, -200))
}
