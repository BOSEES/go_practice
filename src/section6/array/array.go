package main

import "fmt"

func main() {
	// var arr1 [5]int
	// var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	// var arr3 = [5]int{1, 2, 3, 4, 5}
	// arr4 := [5]int{1, 2, 3, 4, 5}
	// arr5 := [5]int{1, 2, 3}
	// arr6 := [...]int{1, 2, 3, 4, 5} //배열크기 자동 맞춤
	// arr7 := [5][5]int{
	// 	{1, 2, 3, 4, 5},
	// 	{1, 2, 3, 4, 5},
	// 	{1, 2, 3, 4, 5},
	// }

	// arr1[2] = 5

	// fmt.Println(arr1)
	// fmt.Println(arr2)
	// fmt.Println(arr3)
	// fmt.Println(arr4)
	// fmt.Println(arr5)
	// fmt.Println(arr6)
	// fmt.Println(arr7)
	// fmt.Printf("%-5T %d %v", arr1, len(arr1), arr1)

	// arr8 := [5]int{1, 2, 3, 4, 5}
	// arr9 := [5]int{
	// 	1,
	// 	2,
	// 	3,
	// 	4,
	// 	5,
	// }
	// arr10 := [...]string{"rla", "dl", "qkr"}
	// fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	// fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
	// fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)

	// arr1 := [5]int{1, 10, 100, 1000, 10000}

	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Println(arr1[i])
	// }

	// for _, value := range arr1 {
	// 	fmt.Println(value)
	// }

	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1

	fmt.Println(&arr1, &arr2)
	fmt.Println(&arr1 == &arr2)

	fmt.Printf("%p %v\n", &arr1, arr1)
	fmt.Printf("%p %v\n", &arr2, arr2)
}
