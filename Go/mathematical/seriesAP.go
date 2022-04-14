package mathematical

import (
	"fmt"
	"strings"
)

func SeriesAP(A1, A2, N int) {
	var sb strings.Builder
	
	jump := A2 - A1
	for i := 0; i < N; i++ {
		val := A1 + (jump * i)
		sb.WriteString(fmt.Sprintf("%-5d", val))
	}

	fmt.Printf("SeriesAP: %s\n", sb.String())
}
