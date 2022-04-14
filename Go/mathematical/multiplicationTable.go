package mathematical

import (
	"fmt"
	"strings"
)

func MultiplicationTable(N int) {

	var sb strings.Builder
	for i := 1; i <= 10; i++ {
		x := i * N
		sb.WriteString(fmt.Sprintf("%-3d", x))
	}
	fmt.Printf("Multiplication Table: %s\n", sb.String())
}
