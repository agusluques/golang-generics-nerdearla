package main

import (
	"fmt"
	"reflect"
)

/*****************************************************
************** BEFORE GENERICS ***********************
******************************************************/

func PrintArrays() {
	arrayOfInts := make([]int, 3)
	arrayOfStrings := make([]string, 3)

	fmt.Printf("len=%d cap=%d %v\n", len(arrayOfInts), cap(arrayOfInts), arrayOfInts)
	fmt.Printf("len=%d cap=%d %v\n", len(arrayOfStrings), cap(arrayOfStrings), arrayOfStrings)
}

func PrintValue(myInterface interface{}) {
	switch v := myInterface.(type) { // HL
	case int:
		// v is an int here, so e.g. v + 1 is possible.
		fmt.Printf("Integer: %v", v)
	case float64:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		fmt.Printf("Float64: %v", v)
	case string:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		fmt.Printf("String: %v", v)
	default:
		// And here I'm feeling dumb. ;)
		fmt.Printf("I don't know, ask stackoverflow.")
	}
}

/*****************************************************
******************************************************
******************************************************/

/*****************************************************
***************** SYNTAX *****************************
******************************************************/

// START SYNTAX1 OMIT
// Define la función f con el parámetro de tipo T
func f[T any](t T) { // HL
	// ...
}

func callF() {
	// Llama a la función f especificando el parámetro int
	f[int](10)                   // HL
	f[string]("hola Nerdearla!") // HL
}

// END SYNTAX1 OMIT

// START SYNTAX2 OMIT
// Define una estructura con parámetro de tipo T
type e[T any] struct { // HL
	t T // HL
}

func useE() {
	// Instancia la estructura especificando el parámetro de tipo int
	valInt := e[int]{t: 1}
	valString := e[string]{t: "hola Nerdearla!"}
	fmt.Println(valInt)
	fmt.Println(valString)
}

// END SYNTAX2 OMIT

// START SYNTAX3 OMIT
// Define una interface con el parámetro de tipo T
type i[T any] interface { // HL
	HacerAlgo(t T) T // HL
}

// END SYNTAX3 OMIT

/*****************************************************
******************************************************
******************************************************/

/*****************************************************
************ TYPE PARAMENTERS FUNC *******************
******************************************************/

func MapStrings(input []string, f func(string) interface{}) []interface{} { // HL
	result := make([]interface{}, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func MapInts(input []int, f func(int) interface{}) []interface{} { // HL
	result := make([]interface{}, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

/*****************************************************
******************************************************
******************************************************/

/*****************************************************
***************** WHEN TO USE ************************
******************************************************/

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// START WHEN1 OMIT
type Tree[T any] struct {
	left, right *Tree[T]
	value       T
}

func (t *Tree[T]) Search(x T) *Tree[T]

var stringTree *Tree[string]

// END WHEN1 OMIT

func GetStringFrom(a interface{}) string {
	reflected := reflect.ValueOf(a)

	switch reflected.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", reflected.Int())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", reflected.Float())
	case reflect.String:
		return reflected.String()
	default:
		return fmt.Sprintf("%v", a)
	}
}

/*****************************************************
******************************************************
******************************************************/
