package combinationSum

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const iterations = 100
const rangeLower = 15
const rangeUpper = 500
const numCandidates = 15

func TestCombinationSum(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < iterations; i++ {
		candidates := make([]int, 0)
		for i := 0; i < numCandidates; i++ {
			randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
			candidates = append(candidates, randomNum)
		}
		target := rangeUpper + rand.Intn(rangeUpper)

		sort.Ints(candidates)

		fmt.Println(candidates, target)

		//CombinationSum2(candidates, target)
		fmt.Println(CombinationSum(candidates, target))
	}
}
