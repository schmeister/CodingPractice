package reverse

func Reverse(input string) string {
	output := []rune(input)
	length := len(output)-1
	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	return string(output)
}
