package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("i: ", i)
	}

	si := []string{"seoul", "busan", "incheon"}

	for index, name := range si {
		fmt.Println(index, name)
	}
}
