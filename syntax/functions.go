package main

import "fmt"

func SumOfInts(l []int) int {
	res := 0
	for _, i := range l {
		res += i
	}
	return res
}

func SumOfFloats(l []float64) float64 {
	res := 0.0
	for _, i := range l {
		res += i
	}
	return res
}

func main() {
	li := []int{1, 2, 3}
	fmt.Println(SumOfInts(li))

	lf := []float64{1, 2, 3}
	fmt.Println(SumOfFloats(lf))
}
