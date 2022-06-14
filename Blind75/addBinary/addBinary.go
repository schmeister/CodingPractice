package addBinary

import "fmt"

func AddBinary(a string, b string) string {

	runeA := []rune(a)
	runeB := []rune(b)
	runeC := make([]rune, 0)

	i := len(a) - 1
	j := len(b) - 1

	carry := rune(0)
	for i >= 0 || j >= 0 {
		va := rune(0)
		vb := rune(0)
		if i >= 0 {
			va = (runeA[i] - '0')
			i--
		}
		if j >= 0 {
			vb = (runeB[j] - '0')
			j--
		}
		bit := (va + vb + carry) % 2
		fmt.Println(va, vb, carry, "=", bit)
		carry = (va + vb + carry) / 2
		runeC = append([]rune{bit + '0'}, runeC...)
	}
	return string(runeC)
}
