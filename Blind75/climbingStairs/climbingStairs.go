package climbingStairs

func ClimbStairs(n int) int {

	// Initialize the climb
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	current := 2
	pre := 1
	//  Fibonacci 1 1 2 3 5 8 13
	for i := 3; i <= n; i++ {
		current = current + pre
		pre = current - pre
	}
	return current
}
