# 总结

* 上面写的的代码大部分不能直接运行，都会 panic，提示“all goroutines are asleep - deadlock!”，因为 go 的 runtime 会检查你所有的 goroutine 都卡住了， 没有一个要执行。

* 可以在阻塞代码前面加上一个或多个你自己业务逻辑的 goroutine，这样就不会 deadlock 了。
