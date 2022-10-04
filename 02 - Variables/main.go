package main

func main() {
	var variavel1 string = "teste"
	println(variavel1)

	variavel2 := "teste2"
	println(variavel2)

	var (
		variavelstrg string = "lalalal"
		variavelint  int    = 10
	)

	println(variavelstrg, variavelint)

	variavelline1, variavelline2, variavelline3 := "teste", 90, "tres"

	println(variavelline1, variavelline2)

	const TESTE_CONST string = "MINHA_CONSTANTE"

	variavelline1, variavelline3 = variavelline3, variavelline1

	println(variavelline1, variavelline3)

}
