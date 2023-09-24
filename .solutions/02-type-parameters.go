package main

// ✅ Ejercicio 1:
func Suma[T int](a, b T) T {
	return a + b
}

// ✅ Ejercicio 2:
func Contiene[T comparable](lista []T, elemento T) bool {
	for _, val := range lista {
		if val == elemento {
			return true
		}
	}
	return false
}
