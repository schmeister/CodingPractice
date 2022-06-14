package longestSubstring

func LongestSubstring(s string) int {

	// more is allocated than needed, but no GC involved
	chars := make(map[byte]int)

	max := 0
	i := 0
	for j := 0; j < len(s); j++ {
		chars[s[j]]++
		for chars[s[j]] > 1 && i < j {
			chars[s[i]]--
			i++
		}

		if max < j-i+1 {
			max = j - i + 1
		}
	}
	return max
}