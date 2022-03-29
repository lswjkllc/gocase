package main

import "fmt"

// 统计小于非负整数 n 的质数的数量
func countPrimes(n int) int {
	// base case
	if n < 3 {
		return 0
	}
	// 初始化一个数组, 默认都是true
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	// 外层: 2 ~ sqrt(n)
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			// 内层: i*i ~ n
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	// 统计
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println("n:1 => 0", countPrimes(1))
	fmt.Println("n:2 => 0", countPrimes(2))
	fmt.Println("n:3 => 1", countPrimes(3))
	fmt.Println("n:10 => 4", countPrimes(10))
}
