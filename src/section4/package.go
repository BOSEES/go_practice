package main

import (
	"fmt"
	// "os"
	"section4/lib"
)

func main() {

	// var name string

	// fmt.Println("이름은??")

	// fmt.Scanf("%s", &name)
	// fmt.Fprintf(os.Stdout, "Hi %s\n", name)

	fmt.Println("숫자가 10보다 큰가?", lib.CheckNum(5))
}
