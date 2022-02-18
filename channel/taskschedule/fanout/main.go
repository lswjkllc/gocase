package main

import (
	"fmt"
	"sync"
	"time"
)

func fanOut(ch <-chan interface{}, outs []chan interface{}, async bool) {
	go func() {
		defer func() { //退出时关闭所有的输出chan
			for i := 0; i < len(outs); i++ {
				close(outs[i])
			}
		}()

		for v := range ch { // 从输入chan中读取数据
			v := v
			for i := 0; i < len(outs); i++ {
				i := i
				if async { //异步
					go func() {
						outs[i] <- v // 放入到输出chan中,异步方式
					}()
				} else {
					outs[i] <- v // 放入到输出chan中，同步方式
				}
			}
		}
	}()
}

func main() {
	// 输入
	in := make(chan interface{})
	go func() {
		defer close(in)
		for i := 0; i < 10; i++ {
			in <- i
		}
	}()

	// 输出
	outs := []chan interface{}{make(chan interface{}), make(chan interface{}), make(chan interface{})}
	fanOut(in, outs, false)

	// 打印输出
	wg := sync.WaitGroup{}
	for i := 0; i < len(outs); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for v := range outs[i] {
				fmt.Printf("%d -> %v\n", i, v)
				time.Sleep(time.Second)
			}
		}(i)
	}
	wg.Wait()
}
