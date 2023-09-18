package main

import "fmt"

func GetStringFrom[T any](a T) string {
	return fmt.Sprintf("%v", a)
}
