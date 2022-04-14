package mathematical

import (
	"fmt"
	"math"
	"strings"
)

func ClosestNumber(N, M int64) {
	var sb strings.Builder

	floor := int64(math.Floor(float64(N)/float64(M))) * M
	ceil := int64(math.Ceil(float64(N)/float64(M))) * M
	difflow := N - floor
	diffhigh := ceil - N
	if difflow < diffhigh {
		sb.WriteString(fmt.Sprintf("%d", floor))
	} else {
		sb.WriteString(fmt.Sprintf("%d", ceil))
	}
	fmt.Printf("ClosestNumber: N=%-5d M=%-5d   %s\n", N, M, sb.String())
}
