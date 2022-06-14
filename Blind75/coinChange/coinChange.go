package coinChange

import (
	"sort"
)

func CoinChange(coins []int, amount int) int {
	sort.Ints(coins)
	numTypes := len(coins) - 1
	count := 0

	for i := numTypes; i >= 0; i-- {
		coinVal := coins[i]
		for coinVal <= amount {
			count++
			amount = amount - coinVal
		}
	}
	if amount != 0 {
		return -1
	}
	return count
}
