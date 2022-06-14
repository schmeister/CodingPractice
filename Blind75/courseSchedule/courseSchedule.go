package courseSchedule

import "fmt"

func CanFinish(n int, pre [][]int) bool {
	depCount := make([]int, n)
	isAReqOf := make([][]int, n)
	for _, v := range pre {
		depCount[v[0]]++
		isAReqOf[v[1]] = append(isAReqOf[v[1]], v[0])
	}
	fmt.Println("depCount:", depCount)
	fmt.Println("idx preReq of:", isAReqOf)

	next := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if depCount[i] == 0 {
			next = append(next, i)
		}
	}
	fmt.Println("Classes with no dependencies:", next)

	for i := 0; i != len(next); i++ {
		class := next[i]
		v := isAReqOf[class]
		fmt.Printf("\t%d Depends on %d \n", v, class)
		for _, vv := range v {
			depCount[vv]--
			fmt.Printf("\tvv %d = %d\n", vv, depCount[vv])
			if depCount[vv] == 0 {
				next = append(next, vv)
			}
		}
	}
	fmt.Println("next", next)
	return len(next) == n
}
