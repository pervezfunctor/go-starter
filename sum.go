package main

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

func Sum[T Number](vals []T) T {
	var s T = 0

	for _, v := range vals {
		s += v
	}

	return s
}
