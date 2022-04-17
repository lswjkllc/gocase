package main

import "fmt"

func insertionSort(nums []int) []int {
	N := len(nums)
	if N < 2 {
		return nums
	}

	for i := 0; i < N; i++ {
		preIdx := i - 1
		current := nums[i]
		for preIdx >= 0 && nums[preIdx] > current {
			nums[preIdx+1] = nums[preIdx]
			preIdx = preIdx - 1
		}
		nums[preIdx+1] = current
	}
	return nums
}

func main() {
	nums := []int{3, 7, 5, 9, 2}
	fmt.Println(insertionSort(nums))
}
