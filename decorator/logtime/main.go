package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	ptr := reflect.ValueOf(i).Pointer()
	rtFunc := runtime.FuncForPC(ptr)
	return rtFunc.Name()
}

func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %v ---\n",
				getFunctionName(f), time.Since(t))
		}(time.Now())
		return f(start, end)
	}
}

func Sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

func main() {
	/*
		1、有两个 Sum 函数，Sum1() 函数就是简单地做个循环，Sum2() 函数动用了数据公式。
		（注意：start 和 end 有可能有负数的情况。）
		2、代码中使用了 Go 语言的反射机制来获取函数名。
		3、修饰器函数是 timedSumFunc()。
	*/

	sum1 := timedSumFunc(Sum1)
	sum2 := timedSumFunc(Sum2)

	fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))
}
