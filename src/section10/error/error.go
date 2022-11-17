package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// func main() {
// 	_, err := os.Open("unnamedfile")

// 	if err != nil {
// 		log.Fatal(err.Error())
// 		fmt.Println("error!!")
// 	}
// 	fmt.Println("========")
// 	fmt.Println("END")
// }

// func isZero(num int) (string, error) {
// 	if num == 0 {
// 		return "", fmt.Errorf("error")
// 	}
// 	s := fmt.Sprint("hello golang", num)
// 	return s, nil
// }

// func main() {
// 	a, err := isZero(12312)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	b, err := isZero(0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// func main() {
// 	var err1 error = errors.New("asdasdsad")
// 	err2 := errors.New("alwdklawkla")

// 	fmt.Println(err1, err2)
// }

// func Power(i float64, a float64) (float64, error) {
// 	if i == 0 {
// 		return i, errors.New("에러입니다")
// 	}
// 	return math.Pow(i, a), nil
// }

// func main() {

// 	if f, err := Power(7, 3); err != nil {
// 		fmt.Printf("Error Message: %s\n", err)
// 	} else {
// 		fmt.Println(f)
// 	}

// 	if f, err := Power(0, 3); err != nil {
// 		fmt.Printf("Error Message: %s\n", err)
// 	} else {
// 		fmt.Println(f)
// 	}
// }

type PowerError struct {
	time    time.Time
	value   interface{}
	message string
}

func (p PowerError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value: %g) - %s", p.time, p.value, p.message)
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowerError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다."}
	}

	if math.IsNaN(f) {
		return 0, PowerError{time: time.Now(), value: f, message: "숫자여야만 합니다."}
	}

	if math.IsNaN(i) {
		return 0, PowerError{time: time.Now(), value: i, message: "숫자여야만 합니다."}
	}

	return math.Pow(f, i), nil
}

func main() {
	v, err := Power(10, 10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)

	t, err := Power(0, 10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)

}
