package main

import (
	"fmt"
	"operations/operations"

	"github.com/badoux/checkmail"
)

func main() {
	println(operations.Soma(1, 2))
	err := checkmail.ValidateFormat("juscilan@gmail.com")
	if err != nil {
		fmt.Println(err)
	}
}
