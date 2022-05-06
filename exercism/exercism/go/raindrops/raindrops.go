package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	t := number % 3
	f := number % 5
	s := number % 7
	div := false
	var sb strings.Builder

	if t == 0 {
		sb.WriteString("Pling")
		div = true
	}
	if f == 0 {
		sb.WriteString("Plang")
		div = true
	}
	if s == 0 {
		sb.WriteString("Plong")
		div = true
	}
	if !div {
		sb.WriteString(strconv.Itoa(number))
	}
	return sb.String()
}
