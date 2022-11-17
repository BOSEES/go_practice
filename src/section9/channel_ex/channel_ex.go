package main

import "fmt"

// func sendOnly(c chan<- int, num int) {
// 	for i := 0; i < 9; i++ {
// 		c <- i
// 	}
// 	c <- 777
// }

// func receiveOnly(c <-chan int) {
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// 	fmt.Println(<-c)
// }

// func main() {
// 	c := make(chan int)

// 	go sendOnly(c, 10)
// 	go receiveOnly(c)

// 	time.Sleep(1 * time.Second)
// }

// func sum(num int) <-chan int {
// 	sum := 0
// 	tot := make(chan int)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			sum += i
// 		}

// 		tot <- sum
// 	}()
// 	return tot
// }

// func main() {
// 	c := sum(100)

// 	fmt.Println(<-c)
// }

func receiveOnly(num int) <-chan int {
	sum := 0
	tot := make(chan int)

	go func() {
		for i := 0; i < num; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 777
		close(tot)
	}()
	return tot
}

func total(c <-chan int) <-chan int {
	tot := make(chan int)

	go func() {
		a := 0

		for i := range c {
			a += i
		}
		tot <- a
	}()

	return tot
}

func main() {
	c := receiveOnly(100)
	output := total(c)

	fmt.Println(<-output)
}
