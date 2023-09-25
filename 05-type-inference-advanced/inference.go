package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// ✏️ Ejercicio 1:
// - Investigar el error de la llamada al método `Print`
// - Hint: chequear cual es el tipo que está infiriendo el compilador
func ScaleAllElements[I constraints.Integer](list []I, factor I) []I {
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
	fmt.Println(ScaleAllElements(listOfInts, 2))

	listOfInts2 := IntSlice{1, 2, 3}
	fmt.Println(ScaleAllElements(listOfInts2, 2))
	listOfInts2.Print()
	ScaleAllElements(listOfInts2, 2).Print()
}
