package main

// ✅ Ejercicio 1:
func Map[T1, T2 interface{}](input []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}
