package mathematical

import (
	"fmt"
	"strings"
)

// for N = 2 the pattern will be
// 2 2 1 1
// 2 1
// for N = 3 the pattern will be
// 3 3 3 2 2 2 1 1 1
// 3 3 2 2 1 1
// 3 2 1
func PrintThePattern(n int) {
	var sb strings.Builder

	for i := n; i > 0; i-- {
		sb.WriteString(fmt.Sprintf("   "))
		for j := n; j > 0; j-- {
			for k := i; k > 0; k-- {
				sb.WriteString(fmt.Sprintf("%d", j))
			}
		}
		sb.WriteString(fmt.Sprintln())
	}
	fmt.Printf("Print The Pattern\n%s",sb.String())
}
