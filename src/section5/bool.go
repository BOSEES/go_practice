package main

import "fmt"

func main() {
	fmt.Println("type 1:", true && true)
	fmt.Println("type 2:", true && false)
	fmt.Println("type 3:", false && true)
	fmt.Println("type 4:", false && false)
	fmt.Println("type 5:", true || true)
	fmt.Println("type 6:", false || true)
	fmt.Println("type 7:", false || false)
	fmt.Println("type 8:", true || false)
}
