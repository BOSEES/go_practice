package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		B
		C
	)

	fmt.Println(A, B, C)
}
