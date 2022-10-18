package main

import "fmt"

func quickSort(nums []int) []int {
	N := len(nums)
	return _quickSort(nums, 0, N-1)
}

func _quickSort(nums []int, left, right int) []int {
	if left < right {
		partitionIdx := partition2(nums, left, right)
		_quickSort(nums, left, partitionIdx-1)
		_quickSort(nums, partitionIdx+1, right)
	}
	return nums
}

func partition2(nums []int, low, high int) int {
	pivot := nums[low]
	if low < high {
		if low < high && nums[high] >= pivot {
			high--
		}
		nums[low] = nums[high]

		if low < high && nums[low] <= pivot {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = pivot
	return low
}

func partition(nums []int, left, right int) int {
	pivot := left
	idx := pivot + 1

	for i := idx; i <= right; i++ {
		if nums[i] < nums[pivot] {
			swap(nums, i, idx)
			idx++
		}
	}
	swap(nums, pivot, idx-1)

	return idx - 1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{3, 7, 5, 9, 2}
	fmt.Println(quickSort(nums))
}
