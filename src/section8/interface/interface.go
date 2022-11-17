package main

import "fmt"

// type test interface{}

// func main() {
// 	var t test
// 	fmt.Println(t)
// }

// type Dog struct {
// 	name   string "개이름"
// 	weight int    "무게"
// }

// func (d Dog) Bite() {
// 	fmt.Println(d.name, "개가 물었다")
// }

// type Behavior interface {
// 	Bite()
// }

// func main() {
// 	dog1 := Dog{"goos", 10}

// 	var interface1 Behavior
// 	interface1 = dog1
// 	interface1.Bite()

// 	dog2 := Dog{"mme", 100}
// 	interface2 := Behavior(dog2)

// 	interface2.Bite()

// 	inter := []Behavior{dog1, dog2}

// 	for index, _ := range inter {
// 		inter[index].Bite()
// 	}

// 	for _, value := range inter {
// 		value.Bite()
// 	}
// }

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func (d Dog) bite() {
	fmt.Println(d.name, "이 물었다")
}
func (d Dog) run() {
	fmt.Println(d.name, "이 뛰다")
}
func (d Dog) sounds() {
	fmt.Println(d.name, "이 울다")
}

//Cat Method

func (d Cat) bite() {
	fmt.Println(d.name, "이 물었다")
}
func (d Cat) run() {
	fmt.Println(d.name, "이 뛰다")
}
func (d Cat) sounds() {
	fmt.Println(d.name, "이 울다")
}

type Behavior interface {
	bite()
	run()
	sounds()
}

func action(a Behavior) {
	a.bite()
	a.run()
	a.sounds()
}

func main() {
	dog := Dog{"alice", 100}
	cat := Cat{"bob", 50}

	action(dog)
	action(cat)
}
