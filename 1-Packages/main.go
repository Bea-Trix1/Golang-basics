package main

import (
	"fmt"
	"module/auxiliary"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing main file")
	auxiliary.Write()
	auxiliary.Write2()
	erro := checkmail.ValidateFormat("beatriz.ferrante@gmail.com")
	fmt.Print(erro)
}
