package leap

// IsLeapYear - is a function
func IsLeapYear(year int) bool {
	// divisible by 4, and not by 100 but odd chance by 400
	return year%4 == 0 && (year%100 != 0 || year%400 == 0) 
}
