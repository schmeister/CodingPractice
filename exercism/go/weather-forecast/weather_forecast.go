// Package weather - This is a package comment.
package weather

// CurrentCondition - condition.
var CurrentCondition string
// CurrentLocation - location.
var CurrentLocation string

// Forecast - yes, the forecast does something.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
