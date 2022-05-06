package hamming

import (
	"fmt"
	"strings"
)

func Distance(a, b string) (int, error) {
	as := strings.Split(a, "")
	bs := strings.Split(b, "")
	if len(as) != len(bs) {
		return 0, fmt.Errorf("mismatch")
	}
	count := 0
	for idx, _ := range as {
		if as[idx] != bs[idx] {
			count++
		}
	}
	return count, nil
}
