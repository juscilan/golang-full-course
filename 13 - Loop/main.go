package main

import "time"

func main() {
	i := 0

	for i < 10 {
		time.Sleep(time.Microsecond)
		i++
	}

	for j := 0; j < 12; j++ {
		println(j)
	}

	nomes := [2]string{"Davi", "lucas"}

	for index, item := range nomes {
		println(index, item)
	}

	for index, item := range "Palavra" {
		println(index, string(item))
	}

	users := map[string]string{
		"Nome":      "Leo",
		"Sobrenome": "Silva",
	}

	for key, value := range users {
		println(key, value)
	}

}
