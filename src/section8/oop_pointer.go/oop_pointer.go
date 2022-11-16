package main

import "fmt"

type shoppingBasket struct {
	cnt   int
	price int
}

func (s shoppingBasket) totalPrice() int {
	return s.cnt + s.price
}

func (s shoppingBasket) copyBasket(cnt, price int) int {
	s.cnt += cnt
	s.price += price

	return s.price + s.cnt
}

func (s *shoppingBasket) refBasket(cnt, price int) int {
	s.cnt += cnt
	s.price += price
	return s.price
}

func main() {

	sp1 := shoppingBasket{10, 1000}

	fmt.Println(sp1.totalPrice())
	fmt.Println(sp1.copyBasket(10, 3000))
	fmt.Println(sp1.totalPrice())
	fmt.Println(sp1.refBasket(1000, 100000))
	fmt.Println(sp1.totalPrice())
}
