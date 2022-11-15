package main

import (
	"fmt"
	"sort"
)

func main() {
	// var slice1 []int
	// slice2 := []int{}
	// slice3 := []int{1, 2, 3, 4, 5}
	// slice4 := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}

	// slice3[3] = 10

	// fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	// fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	// fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	// fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	// slice5 := make([]int, 5, 100)
	// var slice6 []int = make([]int, 5, 100)

	// slice5[2] = 5
	// slice6[1] = 2
	// fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	// fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)

	// slice8 := slice5
	// slice9 := slice3
	// slice9[2] = 1000
	// fmt.Printf("%p %p\n", &slice8, &slice5)
	// fmt.Printf("%p %p %v\n", &slice3, &slice9, slice9)
	// fmt.Printf("%p %p %v\n", &slice3, &slice9, slice3)

	// slice := make([]int, 10, 11)
	// fmt.Printf("%v %p\n", slice, &slice)
	// slice = append(slice, 1)
	// fmt.Printf("%v %p\n", slice, &slice)
	// slice = append(slice, 1)
	// fmt.Printf("%v %p\n", slice, &slice)

	// slice11 := make([]int, 0, 5)

	// for i := 0; i < 15; i++ {
	// 	slice11 = append(slice11, i)
	// 	fmt.Printf("value: %v, length: %d, cap: %d\n", slice11, len(slice11), cap(slice11))
	// }

	slice12 := []int{2, 5, 3, 4, 6, 9, 7, 1}
	slice13 := []string{"a", "c", "f", "z", "w", "h"}

	fmt.Println(sort.IntsAreSorted(slice12))
	sort.Ints(slice12)
	fmt.Println(sort.IntsAreSorted(slice12))
	fmt.Println(slice12)
	sort.Strings(slice13)
	fmt.Println(slice13)

	slice14 := []int{1, 2, 3, 4, 5}
	slice15 := make([]int, 5, 10)
	slice16 := []int{}

	copy(slice15, slice14)
	copy(slice16, slice14)

	fmt.Println(slice15, "\n", slice16)

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)

	b[1] = 10
	fmt.Println(a, "\n", b)

}
