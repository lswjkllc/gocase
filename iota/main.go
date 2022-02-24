package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

/*
左移: x << n ==> x * 2^n
右移: x >> n ==> x / 2^n
*/

const (
	mutextLocked = 1 << iota // 1 << 0 = 1
	mutexWorken              // 1 << 1 = 2
	// mutexStarving                // 1 << 2 = 4
	// mutexWaiterShift = iota // 3      = 3
	mutexWaiterShift = iota // 2      = 2
)

const (
	B  = 1 << (10 * iota) // 1 << (10*0)
	KB                    // 1 << (10*1)
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
	TB                    // 1 << (10*4)
	PB                    // 1 << (10*5)
	EB                    // 1 << (10*6)
	ZB                    // 1 << (10*7)
	YB                    // 1 << (10*8)
)

func main() {
	fmt.Println(mutextLocked, mutexWorken, mutexWaiterShift)
	// fmt.Println(mutextLocked, mutexWorken, mutexStarving, mutexWaiterShift)
	fmt.Println(B, KB, MB, GB, TB, PB, EB)
	var a, b int32 = 0, 0

	go func() {
		atomic.StoreInt32(&a, 1)
		atomic.StoreInt32(&b, 1)
	}()

	for atomic.LoadInt32(&b) == 0 {
		fmt.Println("gosched before...")
		// 让出 CPU 时间片
		runtime.Gosched()
		fmt.Println("gosched after...")
	}
	fmt.Println(atomic.LoadInt32(&a))
}
