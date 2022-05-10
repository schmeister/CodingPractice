package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	r := []rune(input)

	// Short circuit if cannot hold alphabet
	// Works fine without this, but why continue
	if len(r) < 26 {
		return false
	}

	letters := make(map[rune]bool)
	for _, letter := range r {
		if letter < 'a' || letter > 'z' {
			continue
		}
		letters[letter] = true
	}
	if len(letters) < 26 {
		return false
	}
	return true
}
