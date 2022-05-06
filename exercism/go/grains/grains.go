package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		err := errors.New("Invalid number")
		return 0, err
	}

	var v uint64
	v = 1 << (number - 1)

	return v, nil
}

// Total - since this does not take a parameter, just return uint with all 1's in bits.
// It is a constant...
func Total() uint64 {
	var f uint64
	f = 0xFFFFFFFFFFFFFFFF

	return f
}
