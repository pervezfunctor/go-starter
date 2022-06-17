package rational

import (
	"fmt"
)

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

type Rational struct {
	num, den int
}

type Rat interface {
	Num() int
	Den() int
	String() string
	Add(r2 Rat) Rat
	Sub(r2 Rat) Rat
	Mul(r2 Rat) Rat
	Div(r2 Rat) Rat
	Equal(r2 Rat) bool
}

func MakeRational(num, den int) Rat {
	if den == 0 {
		den = 1
	}

	g := Gcd(num, den)
	return Rational{num / g, den / g}
}

func (r Rational) Num() int {
	return r.num
}

func (r Rational) Den() int {
	return r.den
}

func (r Rational) String() string {
	return fmt.Sprintf("%d/%d", r.num, r.den)
}

func (r Rational) Add(r2 Rat) Rat {
	return MakeRational(r.num*r2.Den()+r2.Num()*r.Den(), r.Den()*r2.Den())
}

func (r Rational) Sub(r2 Rat) Rat {
	return MakeRational(r.num*r2.Den()-r2.Num()*r.Den(), r.Den()*r2.Den())
}

func (r Rational) Mul(r2 Rat) Rat {
	return MakeRational(r.num*r2.Num(), r.den*r2.Den())
}

func (r Rational) Div(r2 Rat) Rat {
	return MakeRational(r.num*r2.Den(), r.den*r2.Num())
}

func (r Rational) Equal(r2 Rat) bool {
	return r.num == r2.Num() && r.den == r2.Den()
}
