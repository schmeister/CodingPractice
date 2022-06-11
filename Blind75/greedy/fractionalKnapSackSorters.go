package greedy

// Name sorted
// Len is part of sort.Interface.
func (items NameSortedItems) Len() int {
	return len(items)
}

// Swap is part of sort.Interface.
func (items NameSortedItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (items NameSortedItems) Less(i, j int) bool {
	return items[i].Name < items[j].Name
}

// Len is part of sort.Interface.
func (items ValueSortedItems) Len() int {
	return len(items)
}

// Swap is part of sort.Interface.
func (items ValueSortedItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (items ValueSortedItems) Less(i, j int) bool {
	return items[i].Value < items[j].Value
}

// Len is part of sort.Interface.
func (items WeightSortedItems) Len() int {
	return len(items)
}

// Swap is part of sort.Interface.
func (items WeightSortedItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (items WeightSortedItems) Less(i, j int) bool {
	return items[i].Weight < items[j].Weight
}
// Len is part of sort.Interface.
func (items RatioSortedItems) Len() int {
	return len(items)
}

// Swap is part of sort.Interface.
func (items RatioSortedItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (items RatioSortedItems) Less(i, j int) bool {
	r1 := float64(items[i].Value) / float64(items[i].Weight)
    r2 := float64(items[j].Value) / float64(items[j].Weight)
    return r1 > r2;
}