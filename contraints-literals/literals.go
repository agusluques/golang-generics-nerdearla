package contraintsliterals

/*****************************************************
******************************************************
******************************************************/

// START1 OMIT
func Foo[T int | string, X interface{}](a T, b X) X { // HL
	return b
}

// END1 OMIT

/*****************************************************
******************************************************
******************************************************/

// START2 OMIT
type AllowedValue interface {
	int | string // HL
}

func Bar[T AllowedValue, X interface{}](a T, b X) X { // HL
	return b
}

// END2 OMIT

/*****************************************************
******************************************************
******************************************************/

// START3 OMIT
func Baz[T interface{ int | string }, X interface{}](a T, b X) X { // HL
	return b
}

// END3 OMIT

// START3-any OMIT

// src/builtin/builtin.go

// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}

// END3-any OMIT

/*****************************************************
******************************************************
******************************************************/

// START4 OMIT
func BazBetter[T int | string, X any](a T, b X) X { // HL
	return b
}

// END4 OMIT

/*****************************************************
******************************************************
******************************************************/
