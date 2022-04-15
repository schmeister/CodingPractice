package mathematical

import (
	"fmt"
	"strconv"
	"strings"
)

func Reverse(val int) int {

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
	fmt.Printf("Reverse: %d --> %d\n", val, val2)
	return val2
}
