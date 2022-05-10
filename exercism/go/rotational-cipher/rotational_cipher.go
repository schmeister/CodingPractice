package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	runes := []rune(plain)

	for x := 0; x < len(runes); x++ {
		r := runes[x]
		var add rune = 0
		
		// Set the offset value
		if r >= 'A' && r <= 'Z' {
			add = 'A'
		} else if r >= 'a' && r <= 'z' {
			add = 'a'
		}
		// cipher only if a-z or A-Z
		if add > 0 {
			r = r - add
			for shiftKey < 0 {
				shiftKey += 26
			}

			r = (r + rune(shiftKey))
			r = r % 26
			r += add
			runes[x] = r
		}
	}
	return string(runes)
}
