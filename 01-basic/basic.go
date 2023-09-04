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
