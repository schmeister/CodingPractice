package mathematical

import "fmt"

func gcd(x int, y int) int {
	if x == 0 || y == 0 {
		return 0
	}

	if x == y {
		return y
	}

	if x > y {
		return gcd(x-y, y)
	}
	return gcd(x, y-x)
}

func GCD(x int, y int) {
	fmt.Printf("GCD: %d & %d = %d\n", x, y, gcd(x, y))
}
