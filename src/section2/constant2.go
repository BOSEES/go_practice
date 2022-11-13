package main

import "fmt"

func main() {
	const a, b int = 10, 20
	const (
		c int    = 20
		d int    = 30
		e string = "test"
		f string = "test2"
	)

	fmt.Println(a, b, c, d, e, f)
}
