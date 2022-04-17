package main

import "fmt"

func selectionSort(nums []int) []int {
	N := len(nums)
	if N < 2 {
		return nums
	}

	for i := 0; i < N-1; i++ {
		minIdx := i
		for j := i + 1; j < N; j++ {
			if nums[i] > nums[j] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}

func main() {
	nums := []int{3, 7, 5, 9, 2}
	fmt.Println(selectionSort(nums))
}
