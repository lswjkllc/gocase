package main

import "fmt"

func bubbleSort(nums []int) []int {
	N := len(nums)
	if N < 2 {
		return nums
	}

	for i := 0; i < N-1; i++ {
		for j := 0; j < N-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{3, 7, 5, 9, 2}
	fmt.Println(bubbleSort(nums))
}
