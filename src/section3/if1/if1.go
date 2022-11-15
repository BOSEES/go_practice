package main

import "fmt"

func main() {
	var a int = 20
	b := 15

	if a >= 15 {
		fmt.Println("a는 15 이상")
	} else if b < 15 {
		fmt.Println("b는 15 미만")
	}
}
