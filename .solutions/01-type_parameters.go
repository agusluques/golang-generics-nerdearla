package main

// ✅ Ejercicio 1:
func Map[T interface{}](input []T, f func(T) interface{}) []interface{} {
	result := make([]interface{}, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}
