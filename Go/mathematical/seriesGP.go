package mathematical

import (
	"fmt"
	"math"
	"strings"
)

func SeriesGP(A, R float64, N int) {
	var sb strings.Builder
	var x float64
	for i := 1; i <= N; i++ {
		x = A * math.Pow(R, float64(i-1))
		sb.WriteString(fmt.Sprintf("%-5.0f", x))
	}
	fmt.Printf("SeriesGP: %s\n", sb.String())
}
