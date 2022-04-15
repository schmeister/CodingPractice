package mathematical

import (
	"fmt"
	"strconv"
	"strings"
)

func SumisPalindrome(val int) bool {

	//
	sliceS := strings.Split(strconv.Itoa(val), "")
	sliceI := make([]int, len(sliceS))
	val = 0
	for i, s := range sliceS {
		sliceI[i], _ = strconv.Atoi(s)
		val += sliceI[i]
	}

	sliceS = strings.Split(strconv.Itoa(val), "")
	len := len(sliceS)
	half := len / 2
	for x := 0; x < half; x++ {
		if sliceS[x] != sliceS[len-x-1] {
			fmt.Printf("SumisPalindrome: %d%s%t\n", sliceI, sliceS, false)
			return false
		}
	}
	fmt.Printf("SumisPalindrome: %d%s%t\n", sliceI, sliceS, true)
	return true
}
