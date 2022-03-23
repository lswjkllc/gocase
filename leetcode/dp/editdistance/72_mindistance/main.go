package main

import "fmt"

/*
给你两个单词 word1 和 word2, 请返回将 word1 转换成 word2 所使用的最少操作数。

你可以对一个单词进行如下三种操作:
	插入一个字符
	删除一个字符
	替换一个字符
*/
// https://leetcode-cn.com/problems/edit-distance/

/*
dp[i][j] 表示 word1[0...i] 和 word2[0...j] 的最小编辑距离
*/

// 递归版本
func MinDistance(word1 string, word2 string) int {
	var dp func(int, int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if word1[i] == word2[j] {
			return dp(i-1, j-1) // 什么都不需要做, 此时: dp(i, j) = dp(i-1, j-1)
		}
		return Min(
			// 插入: 直接在 word1[i] 插入一个和 word2[j] 一样的字符
			//       此时 word2[j] 就被匹配了, 前移 j, 继续和 i 对比
			//       操作 +1
			dp(i, j-1)+1,
			// 删除: 直接把 word1[i] 删除, 前移 i, 继续和 j 对比
			//       操作 +1
			dp(i-1, j)+1,
			// 替换: 直接把 word1[i] 替换成 word2[j] 进行匹配
			//       同时前移 i、j 继续对比
			//       操作 +1
			dp(i-1, j-1)+1,
		)
	}
	return dp(len(word1)-1, len(word2)-1)
}

// 使用动态规划进行优化的版本
func minDistance(word1 string, word2 string) int {
	M, N := len(word1), len(word2)

	dp := make([][]int, M+1)
	for i := 0; i <= M; i++ {
		dp[i] = make([]int, N+1)
		dp[i][0] = i
	}
	for j := 1; j <= N; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[M][N]
}

func Min(a ...int) int {
	min := a[0]
	for i := 1; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
	}
	return min
}

func main() {
	fmt.Println("word1:horse word2:ros => 3", minDistance("horse", "ros"))
	fmt.Println("word1:intention word2:execution => 5", minDistance("intention", "execution"))
}
