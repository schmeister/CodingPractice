package diffsquares

func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += (i * i)
	}
	return sum
}

func Difference(n int) int {
	sq := SquareOfSum(n)
	sm := SumOfSquares(n)

	if sq > sm {
		return sq - sm
	}
	return sm - sq
}
