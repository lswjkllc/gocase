package main

import "fmt"

/*
给定两个字符串 text1 和 text2, 返回这两个字符串的最长 公共子序列 的长度。
如果不存在 公共子序列, 返回 0。

一个字符串的 子序列 是指这样一个新的字符串:
	它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如, "ace" 是 "abcde" 的子序列, 但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
*/
// https://leetcode-cn.com/problems/longest-common-subsequence/

// 递归暴力解法
func LongestCommonSubsequence(text1 string, text2 string) int {
	M, N := len(text1), len(text2)

	var dp func(int, int) int
	dp = func(i, j int) int {
		if i == -1 || j == -1 {
			return 0
		}
		if text1[i] == text2[j] {
			return dp(i-1, j-1) + 1
		}
		return Max(dp(i, j-1), dp(i-1, j))
	}

	return dp(M-1, N-1)
}

// 动态规划优化时间复杂度
func longestCommonSubsequence(text1 string, text2 string) int {
	M, N := len(text1), len(text2)
	if M == 0 || N == 0 {
		return 0
	}

	dp := make([][]int, M+1)
	for i := 0; i <= M; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[M][N]
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println("text1:abcde text2:ace => 3", longestCommonSubsequence("abcde", "ace"))
	fmt.Println("text1:abc text2:abc => 3", longestCommonSubsequence("abc", "abc"))
	fmt.Println("text1:abc text2:def => 0", longestCommonSubsequence("abc", "def"))
}
