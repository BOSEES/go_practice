package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- 77777
		}
	}()

	val1, ok1 := <-ch
	fmt.Println("ex: 1", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("ex: 1", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("ex: 1", val3, ok3)

	close(ch)
	val4, ok4 := <-ch
	fmt.Println("ex: 1", val4, ok4)
}
