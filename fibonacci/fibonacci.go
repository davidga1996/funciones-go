package main

import (
	"fmt"
)

func fibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}

}

func main() {
	var input int

	fmt.Print("Número de iteraciones: ")

	fmt.Scan(&input)

	println(fibonacci(input))
}
