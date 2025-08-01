package sorting

import "fmt"
import "strconv"

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
    return fmt.Sprintf("This is the number %.1f",f)
	panic("Please implement DescribeNumber")
}

type NumberBox interface {
	Number() int
}
type NumberInString string

func (strnum NumberInString) Number() int {
    number , _ := strconv.Atoi(string(strnum))
    return number
}
// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
    return fmt.Sprintf("This is a box containing the number %d.0",nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
    if fancyNum, ok := fnb.(FancyNumber); ok {
        val, _ := strconv.Atoi(fancyNum.Value())
		return val
    }
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    if fancyNum, ok := fnb.(FancyNumber); ok {
        val, _ := strconv.Atoi(fancyNum.Value())
        return fmt.Sprintf("This is a fancy box containing the number %d.0",val)
    }
    return "This is a fancy box containing the number 0.0"
}


// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	// A type switch inspects the concrete type of an interface value.
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	case NumberBox:
		return DescribeNumberBox(v)
	default:
		return "Return to sender"
	}
}
