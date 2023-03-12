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

/*
	nuts: 螺母
	bolts: 螺钉
*/
func MatchNutBolt(nuts, bolts []int, left, right int) {
	if left < right {
		// 根据第一个螺母，将螺钉排序
		tmp := nuts[left]
		pivot := sortNums(bolts, left, right, tmp) // pivot 为螺钉的分界点，且 bolts[pivot]=tmp：左边的螺钉都比螺母tmp小，右边的螺钉都比螺母tmp大
		bolts[left], bolts[pivot] = bolts[pivot], bolts[left]
		// 结果：螺钉首元素和螺母首元素匹配

		// 根据第二个螺钉，将螺母排序
		tmp = bolts[left+1]
		pivot = sortNums(nuts, left+1, right, tmp) // pivot 为螺母的分界点，且 nuts[pivot]=tmp：左边的螺母都比螺钉tmp小，右边的螺母都比螺钉tmp大
		nuts[left+1], nuts[pivot] = nuts[pivot], nuts[left+1]
		// 结果：螺母第二元素和螺钉第二元素匹配

		// 汇总：螺母和螺钉 首元素和第二元素 分别匹配

		// 匹配较小部分
		MatchNutBolt(nuts, bolts, left+2, pivot)
		// 匹配较大部分
		MatchNutBolt(nuts, bolts, pivot+1, right)
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
	// nums[left] = target
	return left
}

func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	nuts := []int{5, 3, 2, 7, 1, 6, 4, 9, 8}
	bolts := []int{4, 1, 3, 9, 2, 8, 6, 5, 7}
	MatchNutBolt(nuts, bolts, 0, 8)
	fmt.Println("nuts :", nuts)
	fmt.Println("bolts:", bolts)
	fmt.Printf("nuts == bolts: %v\n", compareSlice(nuts, bolts))
}
