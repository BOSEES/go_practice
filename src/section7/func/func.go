package main

import "fmt"

func main() {
	a := 10
	b := 20
	mul := func(c int) int {
		return a + b + c
	}
	a = 40
	test := mul(30)

	fmt.Println(test)

}
