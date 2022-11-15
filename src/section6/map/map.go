package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map2 := map[string]int{
		"apple":  20,
		"banana": 30,
		"orange": 10,
	}
	fmt.Println(map1)
	fmt.Println(map2)

	fmt.Println(map2["banana"])

	map3 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	for key, value := range map3 {
		fmt.Println(key, value)
	}
}
