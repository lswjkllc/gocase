package main

import "fmt"

func mergeSort(nums []int) []int {
	N := len(nums)
	if N < 2 {
		return nums
	}

	mid := N / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int

	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}

	return result
}

func main() {
	nums := []int{3, 7, 5, 9, 2}
	fmt.Println(mergeSort(nums))
}
