package main

import "golang.org/x/exp/constraints"

// ✅ Ejercicio 1:
func Suma[T int | int32 | int64 | float32 | float64](a, b T) T {
	return a + b
}

// ✅ Ejercicio 2:
func Min[T int | int32 | int64 | float32 | float64 | string](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// ✅ Ejercicio 3:
func GenMin[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
