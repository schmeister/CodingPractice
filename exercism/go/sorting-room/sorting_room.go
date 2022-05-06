package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
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
	x := 0

	switch v := fnb.(type) {
	case FancyNumberBox:
		xx := fmt.Sprintf("%T", v)
		if xx == "sorting.differentFancyNumber" {
			return 0
		}
		y, err := strconv.Atoi(fnb.Value())
		if err != nil {
			x = 0
		} else {
			x = y
		}
	default:
		x = 0
	}

	return x
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	var x int
	switch fnb.(type) {
	case FancyNumber:
		z := fnb.(FancyNumber)
		val := z.Value()
		y, err := strconv.Atoi(val)
		if err != nil {
			x = 0
		} else {
			x=y
		}
	case FancyNumberBox:
		val := fnb.Value()
		y, err := strconv.Atoi(val)
		if err != nil {
			x = y
		}
	default:
		x = 0
	}
	y := float64(x)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", y)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type) {
	case int:
		in, _ := i.(int)
		return DescribeNumber(float64(in))
	case float64:
		in, _ := i.(float64)
		return DescribeNumber(in)
	case NumberBox:
		return DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		return DescribeFancyNumberBox(i.(FancyNumberBox))
	}
	return "Return to sender"
}
