package main

import (
	"encoding/csv"
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
	file, err := os.Create("file_write.csv")
	errCheck1(err)

	defer file.Close()

	wr := csv.NewWriter(file)
	// wr := csv.NewWriter(bufio.newFile)

	wr.Write([]string{"Kim:, 4.8", "park:, 4.8", "lee:, 4.8", "hong:, 4.8"})

	wr.Flush()

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Printf("CSV file write (%d bytes)\n", fi.Size()) // size
	fmt.Println("CSV file name ", fi.Name())
	fmt.Println("CSV file permit", fi.Mode())
}
