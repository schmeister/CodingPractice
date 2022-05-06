package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	details int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.details)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	f, err := weightFodder.FodderAmount()

	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if err == ErrScaleMalfunction && f > 0 {
		return (f * 2 / float64(cows)), nil
	} else if err == ErrScaleMalfunction && f < 0 {
		return 0, errors.New("negative fodder")
	} else if err != nil {
		return 0, err
	}

	if f < 0 {
		return 0, errors.New("negative fodder")
	}

	if cows < 0 {
		return 0, &SillyNephewError{
			details: cows,
		}
	}

	ret := f / float64(cows)
	return ret, nil
}
