package constraints

// ✏️ Ejercicio 1:
// - Generalizar la funcion "Suma" tal que podamos sumar todos los tipos numéricos
// de Go (ints y flotantes)
// - Verificar que los tests pasen
func Suma(a, b int) int {
	return a + b
}

// ✏️ Ejercicio 2:
// - Generalizar la función "Min" tal que podamos utilizarla para todos los tipos
// numéricos y strings
// - Verificar que los tests pasen
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ✏️ Ejercicio 3:
// - Escribir la función "GenMin" usando el paquete constraints para definir el
// type parameter
// - Verificar que los tests pasen
// 💡 Puedes utilizar la intefaz `constraints.Ordered`
func GenMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
