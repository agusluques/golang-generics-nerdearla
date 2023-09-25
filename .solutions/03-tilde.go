package main

import "fmt"

// ✅ Ejercicio 1:
func GetFloatsIn2Decimals[F ~float32 | ~float64](f F) string {
	return fmt.Sprintf("%.2f", f)
}
