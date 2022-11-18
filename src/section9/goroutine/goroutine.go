package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 end")
}
func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 end")
}
func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 end")
}

func main() {
	exe1()
	fmt.Println("main routine start", time.Now())

	go exe2()
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("main routine end")
}
