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
	file, err := os.Create("test_write.txt")
	errCheck1(err)

	defer file.Close()

	s1 := []byte{115, 111, 109, 101, 115}
	n1, err := file.Write(s1)
	errCheck1(err)

	fmt.Printf("쓰기 작업 완료 %d \n", n1)

	file.Sync()

	s2 := "asasas"
	n2, err := file.WriteString(s2)
	errCheck1(err)

	fmt.Printf("Write String to file, %d", n2)

	s3 := "ttttttt"
	n3, err := file.WriteAt([]byte(s3), 70)
	errCheck1(err)

	fmt.Println(n3)

	file.Sync()

	n4, err := file.WriteString("asdasdasasdas")
	errCheck1(err)
	fmt.Println(n4)
}
