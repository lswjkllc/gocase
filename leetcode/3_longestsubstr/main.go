package main

import (
	"fmt"
)

/*
题⽬⼤意
	在⼀个字符串重寻找没有重复字⺟的最⻓⼦串。
解题思路
	这⼀题和第 438 题，第 3 题，第 76 题，第 567 题类似，⽤的思想都是"滑动窗⼝"。
	滑动窗⼝的右边界不断的右移，只要没有重复的字符，就持续向右扩⼤窗⼝边界。
	⼀旦出现了重复字符，就需要缩⼩左边界，直到重复的字符移出了左边界，然后继续移动滑动窗⼝的右边界。
	以此类推，每次移动需要计算当前⻓度，并判断是否需要更新最⼤⻓度，最终最⼤的值就是题⽬中的所求。
*/

func lengthOfLongestSubstring(s string) int {
	// base line
	sLen := len(s)
	if sLen <= 1 {
		return sLen
	}
	// 左右指针 结果
	left, right, res := 0, 0, 0
	// 窗口: 记录每个字符的数量
	window := map[byte]int{}

	for right < sLen {
		c1 := s[right]
		window[c1]++
		right++

		for window[c1] > 1 {
			c2 := s[left]
			window[c2]--
			left++
		}

		res = max(res, right-left)
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("abcabcbb =>", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("bbbbb =>", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("pwwkew =>", lengthOfLongestSubstring("pwwkew"))
}
