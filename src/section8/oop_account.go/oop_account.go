package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	a1 := Account{"1", 100, 0.01}
	a2 := new(Account)
	a2.number = "2"
	a2.balance = 10000
	a2.interest = 0.04
	a3 := &Account{"3", 100, 0.03}

	fmt.Println(a1.Calculate())
	fmt.Println(a2.Calculate())
	fmt.Println(a3.Calculate())
}
