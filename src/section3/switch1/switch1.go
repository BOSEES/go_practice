package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := 7

	switch {
	case a > 0:
		fmt.Println("a is greater 0")
	}

	switch b := 10; {
	case b > 5:
		fmt.Println("b is greater 5")
	}

	switch c := "go"; c {
	case "go":
		fmt.Println("go")
	case "java":
		fmt.Println("java")
	case "python":
		fmt.Println("python")
	default:
		fmt.Println("undefined")
	}

	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("golang")
	}

	switch j, k := 20, 30; {
	case j < k:
		fmt.Println("k is greater j")
	}

	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i is 50~99")
	case i < 50 && i >= 10:
		fmt.Println("i is 10~49")
	case i < 10:
		fmt.Println("i is 0~9")
	}

	switch u := 30 / 15; u {
	case 2, 3, 4:
		fmt.Println(u)
	}

}
