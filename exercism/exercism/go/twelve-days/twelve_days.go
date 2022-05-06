package twelve

import (
	"fmt"
	"strings"
)

var days = []string{"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

var gifts = []string{"a Partridge in a Pear Tree.", "two Turtle Doves,",
	"three French Hens,", "four Calling Birds,",
	"five Gold Rings,", "six Geese-a-Laying,",
	"seven Swans-a-Swimming,", "eight Maids-a-Milking,",
	"nine Ladies Dancing,", "ten Lords-a-Leaping,",
	"eleven Pipers Piping,", "twelve Drummers Drumming,"}

func Verse(i int) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("On the %s day of Christmas my true love gave to me:", days[i-1]))
	for j := i; j > 0; j-- {
		if j == 1 && i > 1 {
			sb.WriteString(" and")
		}
		sb.WriteString(fmt.Sprintf(" %s", gifts[j-1]))
	}
	return sb.String()
}

func Song() string {
	str := ""
	for i := 1; i <= 12; i++ {
		str += Verse(i)
		if i < 12 {
			str += "\n"
		}
	}
	return str
}
