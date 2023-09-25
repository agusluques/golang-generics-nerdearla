package main

import "fmt"

type Rect struct {
	X, Y, W, H int
}

type Point struct {
	X, Y int
}

func GetX[P interface{ Rect | Point }](p P) int {
	return p.X // error ❌: p.X undefined (type P has no field or method X)
}

func main() {
	r := Rect{2, 3, 7, 8}
	p := Point{4, 5}
	fmt.Printf("X: %d %d\n", GetX(r), GetX(p))
}
