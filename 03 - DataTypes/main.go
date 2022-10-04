package main

import (
	"errors"
	"fmt"
)

func main() {
	var inteiro int32 = 8989898
	println(inteiro)

	// int32 alias rune
	var inteiro32 rune = 909909
	println(inteiro32)

	// utint8
	var inteiro8 byte = 7
	println(inteiro8)

	var float32var float32 = 909.09
	println(float32var)

	var float64var float64 = 909.09
	println(float64var)

	char := 'A'
	println(char)

	var boleano1 bool
	println(boleano1)

	var err error = errors.New("Erro")
	fmt.Println(err)
	println(error(err))

}
