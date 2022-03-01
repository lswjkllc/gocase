package main

import "sync"

// 一个已经锁了的 mutex, 再锁一次会一直阻塞。不建议使用
func main() {
	var mu sync.Mutex
	mu.Lock()
	mu.Lock()
}
