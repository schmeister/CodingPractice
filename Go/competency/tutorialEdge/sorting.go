package tutorialEdge

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

// ByAge implements sort.Interface based on the Age field.
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func SortByCustomerComparator() {
	family := []Person{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}
	fmt.Print(family, " --> ") // [{Alice 23} {Eve 2} {Bob 25}]
	sort.Sort(ByAge(family))
	fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}

func SortByStandardLibrary() {
	vals := []int{7, 3, 8, 6, 3}
	fmt.Printf("%d --> ", vals)
	sort.Ints(vals)
	fmt.Printf("%d\n", vals)
}

func SortAMap() {
	m := map[string]int{
		"Eve":   2,
		"David": 2,
		"Bob":   25,
		"Alice": 23}

	fmt.Print(m, " --> ")

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Println(m)
}

func Sorting() {
	SortByStandardLibrary()
	SortByCustomerComparator()
	SortAMap()
}
