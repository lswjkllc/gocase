package main

import (
	"fmt"
	"reflect"
	"time"
)

/// 基于 递归 实现
func fanInRec(chans ...<-chan interface{}) <-chan interface{} {
	switch len(chans) {
	case 0:
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(
			fanInRec(chans[:m]...),
			fanInRec(chans[m:]...))
	}
}

func mergeTwo(a, b <-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for a != nil || b != nil { //只要还有可读的chan
			select {
			case v, ok := <-a:
				if !ok { // a 已关闭，设置为nil
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok { // b 已关闭，设置为nil
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

/// 基于 反射 实现
func fanInReflect(chans ...<-chan interface{}) chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)

		// 构造SelectCase slice
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 循环，从cases中选择一个可用的
		count := 0
		for len(cases) > 0 {
			count += 1
			i, v, ok := reflect.Select(cases)
			if !ok { // 此channel已经close
				fmt.Printf("%d is closed\n", i+count)
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			count -= 1
			out <- v.Interface()
		}
	}()
	return out
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		c <- after
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	// fanin := fanInRec
	fanin := fanInReflect

	out := fanin(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
	)

	for v := range out {
		fmt.Printf("done after %v, output: %v\n", time.Since(start), v)
	}
}
