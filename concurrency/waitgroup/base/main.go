package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("go ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("over")
}
