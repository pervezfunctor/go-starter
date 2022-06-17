package main

import "fmt"

func Swap[T any](x *T, y *T) {
	*x, *y = *y, *x
}

func Id[T any](x T) T {
	return x
}

type Equaler[T any] interface {
	Equal(T) bool
}

type Adder[T any] interface {
	Add(T) T
}

type Suber[T any] interface {
	Sub(T) T
}

type Multer[T any] interface {
	Mul(T) T
}

type Diver[T any] interface {
	Div(T) T
}

type Numeric[T any] interface {
	Equaler[T]
	Adder[T]
	Suber[T]
	Multer[T]
	Diver[T]
	fmt.Stringer
}

func Arith[T Numeric[T]](x, y T) (T, T, T, T) {
	return (x.Add(y)), (x.Sub(y)), (x.Mul(y)), (x.Div(y))
}
