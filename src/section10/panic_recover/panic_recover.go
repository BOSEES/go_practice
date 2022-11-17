package main

import "fmt"

// func main() {
// 	fmt.Println("start")

// 	panic("error user stopped")
// 	fmt.Println("END")
// }

func runFunc() {
	defer func() {
		s := recover()
		fmt.Println("defer", s)
	}()
	panic("user stopped")
}

func main() {
	runFunc()

	fmt.Println("END")
}
