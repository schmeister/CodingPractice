package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	lc := byte('a')
	for i := 0; i < 26; i++ {
		c1 := strings.Count(word, string(lc))
		if c1 >= 2{
			return false
		}
		lc += 1
	}
	return true
}
