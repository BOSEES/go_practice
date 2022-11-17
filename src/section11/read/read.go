package main

import (
	"fmt"
	"os"
)

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	file, err := os.Open("test.txt")
	errCheck1(err)

	fileInfo, err := file.Stat()
	errCheck1(err)

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Size())

	fd1 := make([]byte, fileInfo.Size())
	ct1, err := file.Read(fd1)

	fmt.Printf("%d", ct1)
	fmt.Println(string(fd1))

}
