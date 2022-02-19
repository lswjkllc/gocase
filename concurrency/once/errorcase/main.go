package main

import (
	"fmt"
	"sync"
)

/*
两种错误：
	1、死锁：
		Do 方法会执行一次 f，但是如果 f 中再次调用这个 Once 的 Do 方法的话就会导致死锁。
		这是由于 Lock 的递归调用导致的死锁。
	2、未初始化
*/

func main() {
	var once sync.Once
	once.Do(func() {
		once.Do(func() {
			fmt.Println("初始化")
		})
	})
}
