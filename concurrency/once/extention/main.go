package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

/*
扩展1：CheckOnce
	实现一个类似 Once 的并发原语
	既可以返回当前调用 Do 方法是否正确完成
	还可以在初始化失败后调用 Do 方法再次尝试初始化，直到初始化成功才不再初始化了
*/

// 一个功能更加强大的Once
type CheckOnce struct {
	m    sync.Mutex
	done uint32
}

// 传入的函数f有返回值error，如果初始化失败，需要返回失败的error
// Do方法会把这个error返回给调用者
func (o *CheckOnce) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 { //fast path
		return nil
	}
	return o.slowDo(f)
}

// 如果还没有初始化
func (o *CheckOnce) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 { // 双检查，还没有初始化
		err = f()
		if err == nil { // 初始化成功才将标记置为已初始化
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

/*
扩展2：DoneOnce
	查询 sync.Once 是否执行过
*/
// Once 是一个扩展的sync.Once类型，提供了一个Done方法
type DoneOnce struct {
	sync.Once
}

// Done 返回此Once是否执行过
// 如果执行过则返回true
// 如果没有执行过或者正在执行，返回false
func (o *DoneOnce) Done() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.Once))) == 1
}

func main() {
	var flag DoneOnce
	fmt.Println(flag.Done()) //false

	flag.Do(func() {
		time.Sleep(time.Second)
	})

	fmt.Println(flag.Done()) //true
}
