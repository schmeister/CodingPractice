package validAnagram

func ValidAnagram(str1, str2 string) bool {

	r1 := []rune(str1)
	r2 := []rune(str2)

	m := make(map[rune]int)

	for _, r := range string(r1) {
		count := m[r]
		if count == 0 {
			m[r] = 1
		} else {
			m[r] = m[r] + 1
		}
	}
	for _, r := range string(r2) {
		count := m[r]
		if count == 0 {
			return false
		} else {
			m[r] = m[r] - 1
		}
	}
	for _, r := range m {
		if r != 0 {
			return false
		}
	}
	return true
}
