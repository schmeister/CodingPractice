package robotname

import (
	"errors"
	"fmt"
)

const cLen = 26
const iLen = 10
const capacity int = cLen * cLen * iLen * iLen * iLen

var names map[int]bool = make(map[int]bool, capacity)
var used int = 0
var A rune = 'A'

// Define the Robot type.
type Robot struct {
	sname string
}

func (r *Robot) Name() (string, error) {
	// Robot has no name, give/find one.
	if r.sname == "" {
		if used == capacity {
			return "", errors.New("All used up.")
		}
		r.Reset()
	}
	return r.sname, nil
}
func (r *Robot) Reset() {
	// Initialize the map, if it has not been created..
	if len(names) == 0 {
		for i := 0; i < capacity; i++ {
			names[i] = false
		}
	}

	// Ranging over a map returns in random order...
	var name int
	for k := range names {
		name = k
		break
	}
	delete(names, name)

	letters, numbers := name/1000, name%1000   // top portion for letters, bottom 3 are for numbers.
	letter1, letter2 := letters/26, letters%26 // First letter, second letter
	r.sname = fmt.Sprintf("%s%s%03d", string(A+rune(letter1)), string(A+rune(letter2)), numbers)
	used++
}
