package main

func f[A, B, C any](x A, y B, z C) {}

func main() {
	f(1, true, "go")
	f[int](1, true, "go")
	f[int, bool](1, true, "go")
	f[int, bool, string](1, true, "go")
}
