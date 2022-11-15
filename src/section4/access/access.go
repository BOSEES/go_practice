package main

import (
	check "CheckNum/lib2"
	"fmt"
)

func main() {
	fmt.Println("100 보다 큰 수?", check.CheckNum1(200))
	fmt.Println("1000 보다 큰 수?", check.CheckNum2(700))
}
