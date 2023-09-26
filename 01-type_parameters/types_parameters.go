package typeparameters

// ✏️ Ejercicio 1:
// - Generalizar y definir la función Map con la nueva notación de genéricos
// - Verificar que los tests pasen
func MapStrings(input []string, f func(string) interface{}) []interface{} {
	return Map[string](input, f)
}

func MapInts(input []int, f func(int) interface{}) []interface{} {
	return Map[int](input, f)
}

func Map[T interface{}](input []T, f func(T) interface{}) []interface{} {
	result := make([]interface{}, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}
