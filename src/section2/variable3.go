package main

import "fmt"

func main() {
	//짧은 선언
	shortVar1 := 3
	shortVar2 := "test"
	shortVar3 := false

	fmt.Println("shortVar1: ", shortVar1, "shortVar1: ", shortVar2, "shortVar1: ", shortVar3)

	if i := 10; i < 11 {
		fmt.Println("Short variable Test Success")
	}
}
