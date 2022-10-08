package main

import (
	"fmt"
	"sync"
)

type mouse struct{}

var (
	instance *mouse
	once     = &sync.Once{}
)

func GetInstance() *mouse {
	once.Do(func() {
		instance = new(mouse)
	})
	return instance
}

func main() {
	m := GetInstance()
	fmt.Printf("%T", *m)
}
