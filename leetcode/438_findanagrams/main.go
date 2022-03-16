package main

import "fmt"

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return nil
	}
	// needs
	needs := map[byte]int{}
	for i := 0; i < pLen; i++ {
		needs[p[i]]++
	}
	// window
	window := map[byte]int{}
	// 结果
	res := []int{}
	// 左右指针
	left, right, match := 0, 0, 0

	for right < sLen {
		// 处理右指针
		c1 := s[right]
		if _, ok := needs[c1]; ok {
			window[c1]++
			if window[c1] <= needs[c1] {
				match++
			}
		}
		right++
		// 处理左指针
		for match == pLen {
			if right-left == pLen {
				res = append(res, left)
			}
			c2 := s[left]
			if _, ok := needs[c2]; ok {
				window[c2]--
				if window[c2] < needs[c2] {
					match--
				}
			}
			left++
		}
	}

	return res
}

func main() {
	fmt.Println("cbaebabacd vs abc =>", findAnagrams("cbaebabacd", "abc"))
}
