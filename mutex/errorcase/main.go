package main

import (
	"fmt"
	"sync"
	"time"
)

/*
死锁产生必要条件：
	1、互斥： 至少一个资源是被排他性独享的，其他线程必须处于等待状态，直到资源被释放。
	2、持有和等待：goroutine 持有一个资源，并且还在请求其它 goroutine 持有的资源；
			也就是咱们常说的“吃着碗里，看着锅里”的意思。
	3、不可剥夺：资源只能由持有它的 goroutine 来释放。
	4、环路等待：一般来说，存在一组等待进程，P={P1，P2，…，PN}，P1 等待 P2 持有的资源，P2 等待 P3 持有的资源，依此类推；
			最后是 PN 等待 P1 持有的资源，这就形成了一个环路等待的死结。
避免死锁，只要破坏这四个条件中的一个或者几个，就可以了。
*/

/*
常见的四种 错误场景
*/
func main() {
	/// 第一种：Locak/Unlock 不是成对出现。
	unpaired()
	/// 第二种：Copy 已使用的 Mutex。
	copyUsedMutex()
	/// *** 第三种：重入。Mutex 不是可重入的锁。
	reentrantMutex(&sync.Mutex{})
	/// 第四种：死锁。
	deadlock()
}

func unpaired() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world!")
}

type Counter struct {
	sync.Mutex
	Count int
}

func copyUsedMutex() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo2(c) // 复制锁
}

func foo2(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

func reentrantMutex(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func deadlock() {
	// 派出所证明
	var psCertificate sync.Mutex
	// 物业证明
	var propertyCertificate sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2) // 需要派出所和物业都处理

	// 派出所处理goroutine
	go func() {
		defer wg.Done() // 派出所处理完成

		psCertificate.Lock()
		defer psCertificate.Unlock()

		// 检查材料
		time.Sleep(5 * time.Second)
		// 请求物业的证明
		propertyCertificate.Lock()
		propertyCertificate.Unlock()
	}()

	// 物业处理goroutine
	go func() {
		defer wg.Done() // 物业处理完成

		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()

		// 检查材料
		time.Sleep(5 * time.Second)
		// 请求派出所的证明
		psCertificate.Lock()
		psCertificate.Unlock()
	}()

	wg.Wait()
	fmt.Println("成功完成")
}
