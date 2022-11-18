package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	s := "Goroutine Closure : "

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i)
	}

	time.Sleep(5 * time.Second)
}
