package validPalindrome

import "strings"

func ValidPalindrome(str string) bool {
	str = strings.ToLower(str)

	i := 0
	j := len(str)-1

	for i < j {
		if !isChar(str[i]) {
			i++
			continue
		} else if !isChar(str[j]) {
			j--
			continue
		} else if str[i] != str[j]{
			return false
		}
		i++
		j--
	}
	return true
}

func isChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
