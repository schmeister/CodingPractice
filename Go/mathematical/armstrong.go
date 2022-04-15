package mathematical

import "fmt"

// An Armstrong number of three digits is an integer such
// that the sum of the cubes of its digits is equal to the number itself.
func armstrong(x int) (int, bool) {

	if x < 100 || x > 999 {
		return 0, false
	}

	a := x / 100
	b := (x / 10) % 10
	c := (x) % 10

	val := (a * a * a) + (b * b * b) + (c * c * c)
	if val == x {
		return val, true
	}

	return val, false
}
func RunArmstrong(num int) {
	val, is := armstrong(num)
	fmt.Printf("Armstrong: %3d == %3d = %t\n", num, val, is)
}
