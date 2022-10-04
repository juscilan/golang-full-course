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
}
