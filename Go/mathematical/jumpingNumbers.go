package mathematical

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(val int) int {

	sliceS := strings.Split(strconv.Itoa(val), "")

	len := len(sliceS)
	half := len / 2

	for x := 0; x < half; x++ {
		temp := sliceS[x]
		sliceS[x] = sliceS[len-1-x]
		sliceS[len-1-x] = temp
	}
	str := strings.Join(sliceS, "")
	val2, _ := strconv.Atoi(str)
	return val2
}

func JumpingNumbers(val int) {
	var last int
	for i := 0; i < val; i++ {
		reverse := reverse(i)

		if i == reverse && i > 9 {
			continue
		}

		if i < val {
			last = i
		} else {
			break
		}

		if reverse < val {
			last = reverse
		} else {
			break
		}

	}
	fmt.Printf("JumpingNumbers: %d\n", last)
}
