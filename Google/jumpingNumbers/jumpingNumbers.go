package jumpingNumbers

import (
	"fmt"
	"math"
)

func ListJumpingNumbers(max int) []int {
	values := make([]int, 0)

	for i := 0; i < max; i++ {
		if i <= 10 {
			values = append(values, i)
			continue
		}
		check := 1
		temp := i
		fmt.Printf("%d ", temp)
		digit := temp % 10
		temp /= 10
		fmt.Printf("%d %d\n", temp, digit)

		for temp != 0 {
			if math.Abs(float64(digit-temp%10)) != 1 {
				check = 0
				break
			}
			digit = temp % 10
			temp /= 10
		}
		if check != 0 {
			values = append(values, i)
			fmt.Printf("%d\n", values)
		}
	}
	fmt.Printf("%d\n", values)
	return values
}
