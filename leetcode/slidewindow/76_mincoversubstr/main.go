package main

import (
	"fmt"
)

// 最小覆盖子串
// https://leetcode-cn.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	sLen, tLen := len(s), len(t)
	if sLen < tLen {
		return ""
	}
	// 生成 needs map
	needs := make(map[byte]int, tLen)
	for i := 0; i < tLen; i++ {
		needs[t[i]]++
	}
	// 初始化 window map
	window := make(map[byte]int)

	// 记录最短子串的开始位置和结束位置
	start, end := -1, sLen+1
	// 左右指针
	left, right := 0, 0
	// 匹配度
	match := 0
	for right < sLen {
		// 扩大窗口
		c1 := s[right]
		if _, ok := needs[c1]; ok {
			// 增加有效字符数量
			window[c1]++
			// 只有当 窗口有效字符数量 <= 需要的有效字符数量 时, 增加匹配度
			if window[c1] <= needs[c1] {
				match++
			}
		}
		// right 指针右移
		right++

		for match == tLen {
			// 左右指针长度 < 当前保存的长度时重新设置
			if right-left < end-start {
				start = left
				end = right
			}
			// 缩小窗口
			c2 := s[left]
			if _, ok := needs[c2]; ok {
				// 减少有效字符数量
				window[c2]--
				// 只有当 窗口有效字符数量 < 需要的有效字符数量 时, 减小匹配度
				if window[c2] < needs[c2] {
					match--
				}
			}
			// left 指针右移
			left++
		}
	}
	// 有效性判断
	if start == -1 {
		return ""
	}

	return s[start:end]
}

func main() {
	fmt.Println("ADOBECODEBANC vs ABC =>", minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println("a vs aa =>", minWindow("a", "aa"))
	fmt.Println("a vs a =>", minWindow("a", "a"))
	fmt.Println("a vs b =>", minWindow("a", "b"))
	fmt.Println("aa vs aa =>", minWindow("aa", "aa"))
	fmt.Println("bbaa vs aba =>", minWindow("bbaa", "aba"))
}
