package tilde

type ID int64

func Process[T int64](value T) {
	// ...
}

var id ID = 10
Process(id) // error ❌: ID does not satisfy int64 (possibly missing ~ for int64 in int64)
