package main

func varia(entr ...int) {
	total := 0

	for _, item := range entr {
		total += item
	}

	println(total)
}

func main() {
	varia(1, 2, 2)

}
