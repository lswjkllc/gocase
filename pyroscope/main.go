package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"runtime"
	"sync"

	"github.com/pyroscope-io/client/pyroscope"
)

func cyclenum(num int, wg *sync.WaitGroup) {
	slice := make([]int, num)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			j = i + j
			slice = append(slice, j)
		}
	}
	fmt.Println(len(slice))
	wg.Done()
}

func writeBytes() *bytes.Buffer {
	var buff bytes.Buffer

	for i := 0; i < 300000000; i++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}
	return &buff
}

func run() {
	loop := runtime.GOMAXPROCS(0) * 10
	var wg sync.WaitGroup
	wg.Add(loop)

	for i := 0; i < loop; i++ {
		go cyclenum(300000, &wg)
	}

	writeBytes()

	wg.Wait()
}

func main() {
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "pyroscope-ahhoc-test",

		// replace this with the address of pyroscope server
		ServerAddress: "http://pyroscope-server:4040",

		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,

		// optionally, if authentication is enabled, specify the API key:
		// AuthToken: os.Getenv("PYROSCOPE_AUTH_TOKEN"),

		// by default all profilers are enabled,
		// but you can select the ones you want to use:
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			// pyroscope.ProfileAllocObjects,
			// pyroscope.ProfileAllocSpace,
			// pyroscope.ProfileInuseObjects,
			// pyroscope.ProfileInuseSpace,
		},
	})

	// your code goes here
	run()
}
