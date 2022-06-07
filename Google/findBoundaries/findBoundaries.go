package findBoundaries

import (
	"fmt"
	"math/rand"
	"time"
)

func findLowHigh(key int, arr []int) (int, int) {

	low := findLow(key, arr)
	high := findHigh(key, arr)

	return low, high
}

func findLow(key int, arr []int) int {
	min := 0
	max := len(arr) - 1
	mid := max / 2
	count := 0

	for min <= max {
		fmt.Printf("Low Boundaries: %d %d %d\n", min, mid, max)
		mid_elem := arr[mid]

		if mid_elem < key {
			min = mid + 1
		} else {
			max = mid - 1
		}

		mid = min + (max-min)/2
		count++
	}
	fmt.Println("Iterations:", count)
	if min < len(arr) && arr[min] == key {
		return min
	}

	return -1
}
func findHigh(key int, arr []int) int {
	min := 0
	max := len(arr) - 1
	mid := max / 2
	count := 0

	for min <= max {
		fmt.Printf("High Boundaries: %d %d %d\n", min, mid, max)
		mid_elem := arr[mid]

		if mid_elem <= key {
			min = mid + 1
		} else {
			max = mid - 1
		}

		mid = min + (max-min)/2
		count++
	}

	fmt.Println("Iterations:", count)
	if max == -1 {
		return max
	}
	if max < len(arr) && arr[max] == key {
		return max
	}

	return -1
}

func FindIndex() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 0)
	for x := 0; x < 10; x++ {
		repeats := rand.Intn(35)
		for y := 0; y < repeats; y++ {
			arr = append(arr, x)
		}
	}

arr2 := []int{1,2,2,2,2,2,3,4,5,5,5,5,5,6,7,8,9,9,9,9}

	low, high := findLowHigh(5, arr2)
	if low >= 0 && high >= 0 {
		fmt.Printf("low:%d high:%d\n", low, high)
		fmt.Printf("[%d %d ... %d %d]\n", arr2[low-1], arr2[low], arr2[high], arr2[high+1])
	}
}
