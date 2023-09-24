package basic

import "fmt"

func GetStringFromInt(a int) string {
	return fmt.Sprintf("%d", a)
}

func GetStringFromBoolean(a bool) string {
	return fmt.Sprintf("%v", a)
}

func GetStringFromString(a string) string {
	return a
}

// ✏️ Ejercicio 1:
// - Reescribir esta función utilizando generics
// - Verificar que los tests pasen
func GetStringFrom(a interface{}) string {
	return fmt.Sprintf("%v", a)
}
