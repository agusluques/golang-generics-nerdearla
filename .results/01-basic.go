package results

import "fmt"

func GetString[T any](a T) string {
	return fmt.Sprintf("%v", a)
}
