package main

import "fmt"

func main() {
	const (
		A = iota * 10
		B
		C
		D
	)

	const (
		jan = iota + 1
		feb
		mar
		apr
		may
		jun
	)

	fmt.Println(jan, feb, mar, apr, may, jun)

	fmt.Println(A, B, C, D)
}
