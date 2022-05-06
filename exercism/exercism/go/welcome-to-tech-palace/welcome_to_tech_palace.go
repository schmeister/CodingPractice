package techpalace

import (
	"fmt"
	"regexp"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	output := fmt.Sprintf("%s\n%s\n%s", border, welcomeMsg, border)
	return output
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	m1 := regexp.MustCompile(`\n|\*`)
	str := strings.TrimSpace(m1.ReplaceAllString(oldMsg, ""))

	return str
}
