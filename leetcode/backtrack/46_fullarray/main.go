package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)

	n := len(nums)
	path := make([]int, n)
	copy(path, nums)

	res = backtrack(n, 0, res, nums, path)
	return res
}

func backtrack(n, first int, res [][]int, nums []int, path []int) [][]int {
	if n == first {
		output := make([]int, n)
		copy(output, path)
		res = append(res, output)
		return res
	}
	for i := first; i < n; i++ {
		path[i], path[first] = path[first], path[i]
		res = backtrack(n, first+1, res, nums, path)
		path[i], path[first] = path[first], path[i]
	}
	return res
}

func main() {
	fmt.Println(permute([]int{0, 1}))
}
