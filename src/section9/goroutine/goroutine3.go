package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exe(name int) {
	r := rand.Intn(100)
	fmt.Println(name, " Start:", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>>>>", r, i)
	}
	fmt.Println(name, " End:", time.Now())

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Current system cpu", runtime.GOMAXPROCS(0))

	fmt.Println("Main routine Start :", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("end")
}
