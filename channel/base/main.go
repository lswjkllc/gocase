package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	for i := 1; i <= 4; i++ {
		go func(index int) {
			time.Sleep(time.Duration(index*10) * time.Millisecond)
			for {
				<-ch
				fmt.Printf("I am No %d Goroutine\n", index)
				time.Sleep(time.Second)
				ch <- struct{}{}
			}
		}(i)
	}
	ch <- struct{}{}
	time.Sleep(time.Minute)
}
