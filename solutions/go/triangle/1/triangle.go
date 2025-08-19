package triangle


// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT = iota
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides will return the kind of the given triangle,
func KindFromSides(a, b, c float64) Kind {
	var k Kind
    if a + b > c && a + c > b && b + c > a {
        k = 3
        if a <= 0 || b <= 0 || c <= 0 {
        return 0
    	}
    	switch {
    		case a == b && b == c : return 1
        	case a == b || a == c || b == c : return 2
    	}	
    } 
	return k
}
