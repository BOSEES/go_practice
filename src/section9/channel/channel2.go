package main

import "fmt"

func rangeSum(num int, c chan int) {
	sum := 0

	for i := 1; i < num; i++ {
		sum += i
	}
	c <- sum
}

func main() {
	//channel

	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(700000, c)
	go rangeSum(500000, c)

	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("ex1: ", result1)
	fmt.Println("ex2: ", result2)
	fmt.Println("ex3: ", result3)
}
