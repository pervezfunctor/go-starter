package complex

import "fmt"

type Complex struct {
	real, imag float64
}

func MakeComplex(real, imag float64) Complex {
	return Complex{real, imag}
}

func (c Complex) Real() float64 {
	return c.real
}

func (c Complex) Imag() float64 {
	return c.imag
}

func (c Complex) String() string {
	return fmt.Sprintf("%-.2f + %-.2fi", c.real, c.imag)
}

func Add(c1, c2 Complex) Complex {
	return MakeComplex(c1.real+c2.real, c1.imag+c2.imag)
}

func Sub(c1, c2 Complex) Complex {
	return MakeComplex(c1.real-c2.real, c1.imag-c2.imag)
}

func Mul(c1, c2 Complex) Complex {
	return MakeComplex(c1.real*c2.real-c1.imag*c2.imag, c1.real*c2.imag+c1.imag*c2.real)
}

func Div(c1, c2 Complex) Complex {
	return MakeComplex((c1.real*c2.real+c1.imag*c2.imag)/(c2.real*c2.real+c2.imag*c2.imag), (c1.imag*c2.real-c1.real*c2.imag)/(c2.real*c2.real+c2.imag*c2.imag))
}

func Equal(c1, c2 Complex) bool {
	return c1.real == c2.real && c1.imag == c2.imag
}
