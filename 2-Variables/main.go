package main

import "fmt"

func main() {

	varible := "variable 1"
	var varible2 string = "Variable2"
	fmt.Println(varible, varible2)

	var (
		variable3 string = "ola"
		variable4 string = "mundo"
	)

	fmt.Println(variable3, variable4)

	const constant1 string = "Const1"
	fmt.Println(constant1)

}
