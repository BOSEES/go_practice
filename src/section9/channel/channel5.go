package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
