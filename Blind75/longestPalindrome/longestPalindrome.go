package longestPalindrome

func LongestPalindrome(s string) int {
	// Count individual runes in string
	runeCounter := make(map[rune]int)
	for _, r := range s {
		runeCounter[r]++
	}

	answer := 0
	for _, val := range runeCounter {
		answer += (val / 2) * 2
		if answer%2 == 0 && val%2 == 1 {
			answer++
		}
	}
	return answer
}
