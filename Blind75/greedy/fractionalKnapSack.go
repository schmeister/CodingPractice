package greedy

import (
	"fmt"
	"sort"
)

type Item struct {
	Name   string
	Value  int
	Weight int
}

type NameSortedItems []Item
type ValueSortedItems []Item
type WeightSortedItems []Item
type RatioSortedItems []Item

func StringHeader() string {
	return fmt.Sprintf(" %-10s $%10s %10s\n", "Name", "Value", "Weight")
}
func (i Item) String() string {
	return fmt.Sprintf(" %-10s $%10d %10dkg %.2f\n", i.Name, i.Value, i.Weight, 
	float64(i.Value) / float64(i.Weight))
}

// Main greedy function to solve problem
func fractionalKnapsack(W int, arr []Item, n int) float64 {

	finalvalue := 0.0 // Result (value in Knapsack)

	// Looping through all Items
	for _, val := range arr {
		// If adding Item won't overflow, add it completely
		if val.Weight <= W {
			W -= val.Weight
			finalvalue += float64(val.Value)
		} else {
			finalvalue += (float64(val.Value) * float64(W) / float64(val.Weight))
			break
		}
	}

	// Returning final value
	return finalvalue
}

func FractionalKnapSack() {
	items2 := []Item{
		{
			Value:  60,
			Weight: 10,
		},
		{
			Value:  100,
			Weight: 20,
		},
		{
			Value:  120,
			Weight: 30,
		},
	}
	sort.Sort(RatioSortedItems(items2))
	fmt.Print(StringHeader())
	fmt.Println(items2)
	fmt.Println(fractionalKnapsack(50, items2, 50))
}

func main(){
	FractionalKnapSack()
}

