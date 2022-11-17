package main

import (
	"fmt"
	oper "section12/arithmetic"
)

func main() {
	s := oper.Numbers{100, 10}

	fmt.Println(s.Plus())
	fmt.Println(s.Minus())
	fmt.Println(s.Divide())
	fmt.Println(s.SquareMinus())
	fmt.Println(s.SquarePlus())
}
