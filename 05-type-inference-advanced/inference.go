package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// ✏️ Ejercicio 1:
// - Investigar el error de la llamada al método `Print`
// - Solucionar el error
// - 👀 Hint: chequear cual es el tipo que está infiriendo el compilador
func Scale[I constraints.Integer](list []I, factor I) []I {
	for i := range list {
		list[i] *= factor
	}
	return list
}

type IntSlice []int

func (is IntSlice) Print() {
	fmt.Println(is)
}

func main() {
	listOfInts := []int{1, 2, 3}
	fmt.Println(Scale(listOfInts, 2))
	listOfInts.Print()
}
