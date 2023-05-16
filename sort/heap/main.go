package main

import "fmt"

func swapInts(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func adjustHeap(nums []int, parent, length int) {
	// 保存当前节点的值
	temp := nums[parent]
	// 从当前节点的左子节点开始
	for i := 2*parent + 1; i < length; i = 2*i + 1 {
		// 如果右子节点大于左子节点，则指向右子节点
		if i+1 < length && nums[i] < nums[i+1] {
			i += 1
		}
		// 如果子节点的值大于父节点的值，则将子节点的值赋给父节点（不用进行交换）
		if nums[i] > temp {
			nums[parent] = nums[i]
			parent = i
		} else {
			break
		}
	}
	// 将temp值放到最终的位置
	nums[parent] = temp
}

func heapSort(nums []int) []int {
	N := len(nums)
	if N < 2 {
		return nums
	}
	// 构建大顶堆
	for i := N/2 - 1; i >= 0; i-- {
		// 调整堆
		adjustHeap(nums, i, N)
	}
	fmt.Println(nums)
	// 循环调整堆
	for i := N - 1; i > 0; i-- {
		// 交换堆顶元素与末尾元素
		swapInts(nums, 0, i)
		fmt.Println(nums)
		// 重新调整堆
		adjustHeap(nums, 0, i)
	}

	return nums
}

func main() {
	nums := []int{3, 7, 5, 9, 2}
	fmt.Println("before sort: ", nums)
	fmt.Println("after sort: ", heapSort(nums))
}
