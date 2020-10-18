package main

import (
	"fmt"
)

func plusLarge(numbers ...int) int {
	large := numbers[0]

	for _, v := range numbers {
		if v > large {
			large = v
		}
	}

	return large
}

func main() {
	var count, value int = 0, 0
	var slice []int

	fmt.Print("Número de elementos: ")

	fmt.Scanln(&count)

	for i := 0; i < count; i++ {
		fmt.Scanln(&value)
		slice = append(slice, value)
	}

	fmt.Printf("Número más grande: %d\n", plusLarge(slice...))
}
