package main

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(1, 3)
	println(result)

	f := func(text string) {
		println(text)
	}
	f("funcao f")

	j := func(text string) string {
		println(text)
		return text
	}
	j("funcao f")
	resultado := j("teste")
	println(resultado)
}
