package main

import (
	"fmt"
)

func swap(i, j *int) {
	aux := *i
	*i = *j
	*j = aux
}

func main() {
	var i, j int = 0, 0

	fmt.Print("Ingrese el primer valor:")
	fmt.Scanln(&i)
	fmt.Print("Ingrese el primer valor:")
	fmt.Scanln(&j)

	swap(&i, &j)

	fmt.Printf("Valores intercambiados: %d, %d\n", i, j)
}
