package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums, 判断 nums 中是否存在三个元素 a、b、c,
	使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/
// https://leetcode-cn.com/problems/3sum/

func threeSum(nums []int) [][]int {
	N := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < N; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		twoSum := -nums[i]
		low, high := i+1, N-1
		for low < high {
			tmpSum := nums[low] + nums[high]
			// 记录当前左右指针的值
			left, right := nums[low], nums[high]
			// 三种情况移动左右指针
			if tmpSum < twoSum {
				// 将 左指针 移动至下一个不等于当前值的位置
				for low < high && nums[low] == left {
					low += 1
				}
			} else if tmpSum > twoSum {
				// 将 右指针 移动至下一个不等于当前值的位置
				for low < high && nums[high] == right {
					high -= 1
				}
			} else {
				// 保存满足条件的记录
				ans = append(ans, []int{nums[i], left, right})
				// 将 左指针 移动至下一个不等于当前值的位置
				for low < high && nums[low] == left {
					low += 1
				}
				// 将 右指针 移动至下一个不等于当前值的位置
				for low < high && nums[high] == right {
					high -= 1
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println("nums:[-1,0,1,2,-1,-4] => [[-1,-1,2],[-1,0,1]]", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("nums:[] => []", threeSum([]int{}))
	fmt.Println("nums:[0] => []", threeSum([]int{0}))
}
