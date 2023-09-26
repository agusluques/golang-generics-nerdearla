package main

type Stackeable interface {
	Pop() int
	Push(int)
	Contains(int) bool
}
