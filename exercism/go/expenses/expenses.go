package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	new := make([]Record, 0)
	for _, r := range in {
		b := predicate(r)
		if b {
			new = append(new, r)
		}
	}
	return new
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return c == r.Category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var sum float64
	x := ByDaysPeriod(p)
	recs := Filter(in, x)
	for _, r := range recs {
		sum += r.Amount
	}
	return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {

	// Get Category filter func, filter records, and count records.
	cat := ByCategory(c)
	recs := Filter(in, cat)
	if len(recs) == 0 {
		err := errors.New(fmt.Sprintf("unknown category %s", c))
		return 0, err
	}

	sum := TotalByPeriod(recs, p)
	return sum, nil
}