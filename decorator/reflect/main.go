package main

import (
	"fmt"
	"reflect"
)

/*
用了 reflect.MakeFunc() 函数制作出了一个新的函数，其中的 targetFunc.Call(in) 调用了被修饰的函数。
关于 Go 语言的反射机制，推荐官方文章——《The Laws of Reflection》。
Decorator() 需要两个参数：第一个是出参 decoPtr ，就是完成修饰后的函数。第二个是入参 fn ，就是需要修饰的函数。
*/
func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}

func foo(a, b, c int) int {
	return a + b + c
}

func bar(a, b int) int {
	return a * b
}

func main() {
	var decoFoo func(int, int, int) int
	Decorator(&decoFoo, foo)
	decoFoo(1, 2, 3)

	var decoBrr func(int, int) int
	Decorator(&decoBrr, bar)
	decoBrr(4, 5)
}
