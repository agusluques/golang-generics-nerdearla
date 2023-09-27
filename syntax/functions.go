package main

import "fmt"

func GetStringFrom(s interface{}) string {
	return fmt.Sprintf("%v", s)
}
