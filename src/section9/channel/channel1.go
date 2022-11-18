package main

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("Work1 : Start ->>>", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : End ->>>", time.Now())

	v <- 1
}

func work2(v chan int) {
	fmt.Println("Work2 : Start ->>>", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : End ->>>", time.Now())

	v <- 2
}

func main() {
	fmt.Println("Main: Start - >>", time.Now())

	// var c chan int
	// c = make(chan int)

	v := make(chan int)

	go work1(v)
	go work2(v)

	<-v
	<-v
	fmt.Println("Main: End - >>", time.Now())
}
