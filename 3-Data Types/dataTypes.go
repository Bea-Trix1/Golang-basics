package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int16 = 100
	fmt.Println(number)

	var number2 int64 = 1000000000000000
	fmt.Println(number2)

	var number3 int32 = 1000000000
	fmt.Println(number3)

	var number4 int8 = 100
	fmt.Println(number4)

	//int32 = rune
	var number5 rune = 12345
	fmt.Println(number5)

	// real numbers

	var n1 float32 = 123.32
	var n2 float64 = 123.12312
	fmt.Println(n1, n2)

	//Strings
	var str string = "Hi"
	fmt.Println(str)

	//Boolean
	var boolean1 bool
	fmt.Println(boolean1)

	var erro error = errors.New("Intern error")
	fmt.Println(erro)
}
