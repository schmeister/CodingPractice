package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind = NaT

	if a+b <= c || a+c <= b || b+c <= a {
		// just a point in space (or negative value), not a triangle
		k = NaT
	} else if a == b && a == c {
		k = Equ // equilateral
	} else if a == b || a == c || b == c {
		k = Iso // isosceles
	} else {
		k = Sca // scalene
	}
	return k
}
