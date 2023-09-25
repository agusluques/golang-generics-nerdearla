package typeparameters

// ✏️ Ejercicio 1:
// - Generalizar y definir la función Map con la nueva notación de genéricos
// - Verificar que los tests pasen
func MapStrings(input []string, f func(string) interface{}) []interface{} {
	result := make([]interface{}, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func MapInts(input []int, f func(int) interface{}) []interface{} {
	result := make([]interface{}, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}
