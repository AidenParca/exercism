package complexnumbers

import "math"

// Number represents a complex number in the form a + b*i.
// The provided struct in the prompt had an extra field 'c', which is not needed for a standard complex number.
// It has been removed for this implementation.
type Number struct {
	a, b float64
}

// Real returns the real part 'a' of the complex number.
func (n Number) Real() float64 {
	return n.a
}

// Imaginary returns the imaginary part 'b' of the complex number.
func (n Number) Imaginary() float64 {
	return n.b
}

// Add sums two complex numbers and returns the result.
// (a + i*b) + (c + i*d) = (a + c) + (b + d)*i
func (n1 Number) Add(n2 Number) Number {
	return Number{a: n1.a + n2.a, b: n1.b + n2.b}
}

// Subtract finds the difference between two complex numbers.
// (a + i*b) - (c + i*d) = (a - c) + (b - d)*i
func (n1 Number) Subtract(n2 Number) Number {
	return Number{a: n1.a - n2.a, b: n1.b - n2.b}
}

// Multiply calculates the product of two complex numbers.
// (a + i*b) * (c + i*d) = (a*c - b*d) + (b*c + a*d)*i
func (n1 Number) Multiply(n2 Number) Number {
	realPart := n1.a*n2.a - n1.b*n2.b
	imaginaryPart := n1.b*n2.a + n1.a*n2.b
	return Number{a: realPart, b: imaginaryPart}
}

// Times multiplies a complex number by a real-valued scalar factor.
func (n Number) Times(factor float64) Number {
	return Number{a: n.a * factor, b: n.b * factor}
}

// Divide divides one complex number by another.
// (a + i*b) / (c + i*d) = (a*c + b*d)/(c^2 + d^2) + (b*c - a*d)/(c^2 + d^2)*i
func (n1 Number) Divide(n2 Number) Number {
	denominator := n2.a*n2.a + n2.b*n2.b
	realPart := (n1.a*n2.a + n1.b*n2.b) / denominator
	imaginaryPart := (n1.b*n2.a - n1.a*n2.b) / denominator
	return Number{a: realPart, b: imaginaryPart}
}

// Conjugate returns the complex conjugate of a number.
// The conjugate of (a + i*b) is (a - i*b).
func (n Number) Conjugate() Number {
	return Number{a: n.a, b: -n.b}
}

// Abs calculates the absolute value (modulus) of a complex number.
// |z| = sqrt(a^2 + b^2)
func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

// Exp calculates the exponential of a complex number.
// e^(a + i*b) = e^a * (cos(b) + i*sin(b))
func (n Number) Exp() Number {
	expA := math.Exp(n.a)
	realPart := expA * math.Cos(n.b)
	imaginaryPart := expA * math.Sin(n.b)
	return Number{a: realPart, b: imaginaryPart}
}