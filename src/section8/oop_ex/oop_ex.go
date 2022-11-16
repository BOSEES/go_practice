package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account { // 포인터를 참조해서 리턴 , 원본
	return &Account{number, balance, interest}
}

// func main() {
// 	//구조체 생성자 패턴
// 	kim := Account{number: "245-200", balance: 1000, interest: 0.2}
// 	var lee *Account = new(Account)
// 	lee.balance = 1000
// 	lee.interest = 0.2
// 	lee.number = "200-100"

// 	fmt.Println(kim)
// 	fmt.Println(lee)

// 	park := NewAccount("245-200", 1000, 0.2)

// 	fmt.Println(park)
// }

// func CalculateD(a Account) {
// 	a.balance = a.balance + (a.balance * a.interest)
// }

// func CalculateP(a *Account) {
// 	a.balance = a.balance + (a.balance * a.interest)
// }

// func main() {
// 	kim := Account{"245-900", 100, 0.01}
// 	lee := Account{"245-910", 100, 0.02}

// 	fmt.Println(kim)
// 	fmt.Println(lee)

// 	CalculateD(kim)
// 	CalculateP(&lee)

// 	fmt.Println(kim.balance)
// 	fmt.Println(lee.balance)
// }

// func (a Account) CalculateD(bonus float64) {
// 	a.balance = a.balance + (a.balance * a.interest) + bonus
// }

// func (a *Account) CalculateP(bonus float64) {
// 	a.balance = a.balance + (a.balance * a.interest) + bonus
// }

// func main() {
// 	kim := Account{"245-900", 100, 0.01}
// 	lee := Account{"245-910", 100, 0.02}

// 	kim.CalculateD(100)
// 	lee.CalculateP(200)

// 	fmt.Println(kim.balance)
// 	fmt.Println(lee.balance)
// }

// type Employee struct {
// 	name   string
// 	salary float64
// 	bonus  float64
// }

// func (e Employee) Calculate() float64 {
// 	return e.salary + e.bonus
// }

// type Executives struct {
// 	Employee
// 	specialBonus float64
// }

// func main() {
// 	em1 := Employee{"alice", 1000, 0.1}
// 	em2 := Employee{"bob", 2000, 0.1}

// 	fmt.Println(em1.Calculate())
// 	fmt.Println(em2.Calculate())

// 	ex1 := Executives{
// 		Employee{"jack", 10000, 100},
// 		1000000,
// 	}

// 	fmt.Println(ex1.Calculate())
// }

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee
	specialBonus float64
}

func main() {
	em1 := Employee{"alice", 1000, 0.1}
	em2 := Employee{"bob", 2000, 0.1}

	fmt.Println(em1.Calculate())
	fmt.Println(em2.Calculate())

	ex1 := Executives{
		Employee{"jack", 10000, 100},
		1000000,
	}

	fmt.Println(ex1.Calculate())
	fmt.Println(ex1.Employee.Calculate())
}
