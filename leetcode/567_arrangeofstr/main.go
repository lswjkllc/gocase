package main

import "fmt"

/*
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。
如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。
*/
// https://leetcode-cn.com/problems/permutation-in-string/

func checkInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	needs := map[byte]int{}
	for i := 0; i < s1Len; i++ {
		needs[s1[i]]++
	}
	// 左右指针
	left, right := 0, 0
	// 滑动窗口 和 match
	window, match := map[byte]int{}, 0

	for right < s2Len {
		c1 := s2[right]
		if _, ok := needs[c1]; ok {
			window[c1]++
			if window[c1] <= needs[c1] {
				match++
			}
		}
		right++

		for match == s1Len {
			if right-left == s1Len {
				return true
			}
			c2 := s2[left]
			if _, ok := needs[c2]; ok {
				window[c2]--
				if window[c2] < needs[c2] {
					match--
				}
			}
			left++
		}
	}

	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}
