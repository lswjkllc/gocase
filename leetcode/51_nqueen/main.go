package main

import "fmt"

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	path := make([][]byte, n)
	for i := 0; i < n; i++ {
		path[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			path[i][j] = '.'
		}
	}

	res = backtrack(n, res, path, 0)

	return res
}

func backtrack(n int, res [][]string, path [][]byte, row int) [][]string {
	// 触发结束条件
	if n == row {
		q := make([]string, n)
		p := make([][]byte, n)
		copy(p, path)
		for i := 0; i < n; i++ {
			q[i] = string(p[i])
		}
		res = append(res, q)
		return res
	}

	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValidPath(n, path, row, col) {
			continue
		}
		// 做选择
		path[row][col] = 'Q'
		// 进入下一行决策
		res = backtrack(n, res, path, row+1)
		// 撤销选择
		path[row][col] = '.'
	}

	return res
}

func isValidPath(n int, path [][]byte, row, col int) bool {
	for i := 0; i < row; i++ {
		// 检查 col 列是否已经存在 Q
		if path[i][col] == 'Q' {
			return false
		}
	}
	// 检查 左上 是否存在 Q
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if path[i][j] == 'Q' {
			return false
		}
	}
	// 检查 右上 是否存在 Q
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if path[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solveNQueens(5))
}
