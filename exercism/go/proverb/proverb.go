package proverb

import "fmt"

// Proverb - proverb.
func Proverb(rhyme []string) []string {

	// Not enough wants
	if len(rhyme) < 1 {
		return []string{}
	}

	// Proverb lines
	line := "For want of a %s the %s was lost."
	end := "And all for the want of a %s."

	// Make the proverb
	result := make([]string, 0)
	for x := 0; x < (len(rhyme) - 1); x++ {
		result = append(result, fmt.Sprintf(line, rhyme[x], rhyme[x+1]))
	}
	result = append(result, fmt.Sprintf(end, rhyme[0]))
	return result
}
