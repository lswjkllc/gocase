package main

import "fmt"

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。

===> 这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
===> 同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
*/
// https://leetcode-cn.com/problems/house-robber-ii/

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// 相对 第 196 题 新增约束: 首尾房间相连, 不能同时被抢
	// 情形一: 掐头去尾 ---> 首尾房间都不抢
	// 情形二: 去尾    ---> 第一间房子被抢最后一间不抢
	// 情形三: 掐头    ---> 最后一间房子被抢第一间不抢
	// 只要比较 情况二和情况三 就行了，因为这两种情况对于房子的选择余地比 情况一 大
	// 房子里的钱数都是非负数，所以选择余地大，最优决策结果肯定不会小。
	return Max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

// 仅计算闭区间 [start,end] 的最优结果
func robRange(nums []int, start, end int) int {
	dp_i_1, dp_i_2 := 0, 0
	dp_i := 0
	for i := end; i >= start; i-- {
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
	fmt.Println("nums:[2,3,2] => 3", rob([]int{2, 3, 2}))
	fmt.Println("nums:[1,2,3,1] => 4", rob([]int{1, 2, 3, 1}))
	fmt.Println("nums:[1,2,3] => 3", rob([]int{1, 2, 3}))
	fmt.Println("nums:[1] => 1", rob([]int{1}))
}
