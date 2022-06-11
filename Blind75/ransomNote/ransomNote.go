package ransomNote

import (
	"strings"
)

func CanConstruct(ransomNote string, magazine string) bool {

	m := make(map[rune]int)

	ransomNote = strings.ToLower(ransomNote)
	magazine = strings.ToLower(magazine)

	rr := []rune(ransomNote)
	mr := []rune(magazine)

	for _, val := range rr {
		if isChar(val) {
			num := m[val]
			m[val] = num + 1
		}
	}
	for _, val := range mr {
		if isChar(val) {
			num := m[val]
			m[val] = num - 1
		}
	}
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}

func isChar(r rune) bool {
	return r >= 'a' && r <= 'z'
}
