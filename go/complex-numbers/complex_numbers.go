package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	real, imag float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imag
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.real + n2.real, n1.imag + n2.imag}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.real - n2.real, n1.imag - n2.imag}
}

func (n1 Number) Multiply(n2 Number) Number {
	real := n1.real*n2.real - n1.imag*n2.imag
	imag := n1.real*n2.imag + n1.imag*n2.real
	return Number{real, imag}
}

func (n Number) Times(factor float64) Number {
	return Number{n.real * factor, n.imag * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	denom := n2.real*n2.real + n2.imag*n2.imag
	real := (n1.real*n2.real + n1.imag*n2.imag) / denom
	imag := (n1.imag*n2.real - n1.real*n2.imag) / denom
	return Number{real, imag}
}

func (n Number) Conjugate() Number {
	return Number{n.real, -n.imag}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imag*n.imag)
}

func (n Number) Exp() Number {
	expReal := math.Exp(n.real)
	return Number{
		expReal * math.Cos(n.imag),
		expReal * math.Sin(n.imag),
	}
}
