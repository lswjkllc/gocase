package main

import "fmt"

/*
螺母 nut
螺钉 bolt
匹配无序的螺母和螺钉
规则:
	螺母和螺母之间无法比较
	螺钉和螺钉之间也无法比较
	螺母和螺钉之间可以比较
*/

// 参考链接: https://blog.csdn.net/summer2day/article/details/95979090

func MatchNutBolt(nuts, bolts []int, left, right int) {
	if left < right {
		// 根据第一个螺母，将螺钉排序
		tmp := nuts[left]
		pivotIdx := sortNums(bolts, left, right, tmp)
		bolts[left], bolts[pivotIdx] = bolts[pivotIdx], bolts[left]
		// 结果：螺钉首元素和螺母首元素匹配

		// 根据第二个螺钉，将螺母排序
		tmp = bolts[left+1]
		pivotIdx = sortNums(nuts, left+1, right, tmp)
		nuts[left+1], nuts[pivotIdx] = nuts[pivotIdx], nuts[left+1]
		// 结果：螺母第二元素和螺钉第二元素匹配

		// 汇总：螺母和螺钉 首元素和第二元素 分别匹配

		// 匹配较小部分
		MatchNutBolt(nuts, bolts, left+2, pivotIdx)
		// 匹配较大部分
		MatchNutBolt(nuts, bolts, pivotIdx+1, right)
	}
}

func sortNums(nums []int, left, right int, target int) int {
	for left < right {
		for left < right && nums[left] < target {
			left++
		}
		for left < right && nums[right] > target {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[left] = target
	return left
}

func main() {
	nuts := []int{5, 3, 2, 7, 1, 6, 4, 9, 8}
	bolts := []int{4, 1, 3, 9, 2, 8, 6, 5, 7}
	MatchNutBolt(nuts, bolts, 0, 8)
	fmt.Println("nuts :", nuts)
	fmt.Println("bolts:", bolts)
}
