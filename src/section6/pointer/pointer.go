package main

import "fmt"

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {
	var a *int
	b := 8

	a = &b
	c := &b
	d := &b

	fmt.Println(a, *a, &a)
	fmt.Println(*c, *d)

	*a = 20
	fmt.Println(*c, *d)

	rptc(&b)
	vptc(b)

	fmt.Println(a, b)
}
