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

	for i := n; i > 0; i-- {
		var sb strings.Builder
		for j := n; j > 0; j-- {
			for k := i; k > 0; k-- {
				sb.WriteString(fmt.Sprintf("%d", j))
			}
		}
		fmt.Printf("%s\n", fmt.Sprintf("Print The Pattern: %s", sb.String()))
	}
}
