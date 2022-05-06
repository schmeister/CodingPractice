package logs

import (
	"fmt"
	"strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	if strings.ContainsRune(log, '❗') {
		return "recommendation"
	}
	if strings.ContainsRune(log, '🔍') {
		return "search"
	}
	if strings.ContainsRune(log, '☀') {
		return "weather"
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for idx, _ := range runes {
		if runes[idx] == oldRune {
			runes[idx] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runes := []rune(log)
	fmt.Println(runes, len(runes) , limit)
	if len(runes) <= limit {
		return true
	}
	return false
}
