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

// partition3 vs partition2:
// "左右指针 都移动完 之后再交换" vs "左右指针 移动之后 分别交换" 更容易理解
func partition3(nums []int, left, right int) int {
	// 基准索引初始化为左指针
	pivotIdx := left
	// 外循环
	for left < right {
		// 右指针: 从右向左扫描, 找到 第一个一个小于基准值 的索引
		for left < right && nums[right] >= nums[pivotIdx] {
			right--
		}
		// 左指针: 从左向右扫描, 找到 第一个大于基准值 的索引
		for left < right && nums[left] <= nums[pivotIdx] {
			left++
		}
		// 左右指针索引不同时, 交换左右指针的值
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	// 基准索引和左指针不相等，表示需要交换
	if pivotIdx != left {
		nums[pivotIdx], nums[left] = nums[left], nums[pivotIdx]
	}
	// 响应左指针（此时, 左右指针理论上是相等的）
	return left
}

func partition2(nums []int, left, right int) int {
	// 将左指针的值选定为 基准值
	pivot := nums[left] //! 第一次 保存左指针指针的值 -------
	// 外循环
	for left < right {
		// 右指针: 从右向左扫描, 找到 第一个一个小于基准值 的索引
		for left < right && nums[right] >= pivot {
			right--
		}
		// 将 左指针的值 替换为 右指针的值
		nums[left] = nums[right] //! 保存右指针的值 ------- 原左指针的值保存在 基准/右指针 值中

		// 左指针: 从左向右扫描, 找到 第一个大于基准值 的索引
		for left < right && nums[left] <= pivot {
			left++
		}
		// 将 右指针的值 替换为 左指针的值
		nums[right] = nums[left] //! 保存左指针的值 ------- 原右指针的值保存在左指针中
	}
	// 将 左指针的值 替换为 基准值
	nums[left] = pivot //! 保存基准值（此时, left == right, left中的值和上一轮中的right是重复的）
	// 响应左指针（此时, 左右指针理论上是相等的）
	return left
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
