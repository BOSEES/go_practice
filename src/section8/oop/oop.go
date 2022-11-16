package main

import "fmt"

type Car struct {
	name  string
	color string
	price uint32
	tax   uint32
}

func (c Car) Price() uint32 {
	return c.price + c.tax
}

func (c Car) Color() string {
	return c.color
}

func main() {
	bmw := Car{name: "520d", color: "white", price: 500000, tax: 50000}

	fmt.Println(bmw.Price())
	fmt.Println(bmw.Color())
}
