package main

import "golang.org/x/exp/constraints"

func ScaleAllElements[I constraints.Integer, R ~[]I](list R, factor I) R {
	for i := range list {
		list[i] *= factor
	}
	return list
}
