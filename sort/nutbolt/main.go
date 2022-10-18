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
		i, j := left, right
		for i < j {
			for i < j && bolts[i] < tmp {
				i++
			}
			for i < j && bolts[j] > tmp {
				j--
			}
			if i < j {
				bolts[i], bolts[j] = bolts[j], bolts[i]
			}
		}
		bolts[i] = tmp
		bolts[left], bolts[i] = bolts[i], bolts[left]
		// 结果：螺钉首元素和螺母首元素匹配

		// 根据第二个螺钉，将螺母排序
		tmp = bolts[left+1]
		i, j = left+1, right
		for i < j {
			for i < j && nuts[i] < tmp {
				i++
			}
			for i < j && nuts[j] > tmp {
				j--
			}
			if i < j {
				nuts[i], nuts[j] = nuts[j], nuts[i]
			}
		}
		nuts[i] = tmp
		nuts[left+1], nuts[i] = nuts[i], nuts[left+1]
		// 结果：螺母第二元素和螺钉第二元素匹配

		// 汇总：螺母和螺钉 首元素和第二元素 分别匹配

		// 匹配较小部分
		MatchNutBolt(nuts, bolts, left+2, i)
		// 匹配较大部分
		MatchNutBolt(nuts, bolts, i+1, right)
	}
}

func main() {
	nuts := []int{5, 3, 2, 7, 1, 6, 4, 9, 8}
	bolts := []int{4, 1, 3, 9, 2, 8, 6, 5, 7}
	MatchNutBolt(nuts, bolts, 0, 8)
	fmt.Println("nuts :", nuts)
	fmt.Println("bolts:", bolts)
}
