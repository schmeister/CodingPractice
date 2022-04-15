package mathematical

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func KthDigit(A, B float64, K int) {

	pow := math.Pow(A, B)
	powS := strconv.Itoa(int(pow))
	powStr := strings.Split(powS, "")
	len := len(powStr)
	//	fmt.Println(powStr, len)

	if K > len {
		panic("KthDigit")
	}
	digit, _ := strconv.Atoi(powStr[len-K])
	fmt.Printf("KthDigit: %.0f^%.0f = %.0f  (%d = %d)\n", A,B, pow, K, digit)

}
