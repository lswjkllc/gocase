package main

import (
	"fmt"
	"sync"
)

type TestStruct struct {
	Wait sync.WaitGroup
}

func main() {
	w := sync.WaitGroup{}
	w.Add(1)
	t := &TestStruct{
		Wait: w, // Error: Copy WaitGroup
	}
	fmt.Printf("%p %p\n\n", &w, &t.Wait)

	t.Wait.Done()
	fmt.Println("Finished")
}
