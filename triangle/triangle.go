package triangle

// Required for using IsNaN function
import "math"

func KindFromSides(a, b, c float64) Kind {
    // First part of the function does the mandatory tests
    if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b <= c || a+c <= b || b+c <= a {
		return NaT
	}

    // If triangle is possible, second part identifies the kind
	if a == b && a == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	} else {
		return Sca
	}
}

// Just an int shall suffice
type Kind int

// Just learned about the existence of 'iota', good tool for enums
const (
	Equ Kind = iota // equilateral
	Iso Kind = iota // isosceles
	Sca Kind = iota // scalene
	NaT Kind = iota // not a triangle
)
