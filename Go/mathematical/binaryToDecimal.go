package mathematical

import (
	"fmt"
	"math"
	"strconv"
)

func BinaryToDecimal(binary string) {

	// Simple solution - use built in function
	dec, _ := strconv.ParseInt(binary, 2, 0)

	// Harder solution manual conversion
	dec2, _ := strconv.Atoi(binary)
	var result int
	var position float64
	for dec2 != 0 {
		remainder := dec2 % 10
		result += remainder * int(math.Pow(2, position))
		dec2 /= 10
		position++
	}
	fmt.Printf("BinaryToDecimal: %s %d %d\n", binary, dec, result)
}
