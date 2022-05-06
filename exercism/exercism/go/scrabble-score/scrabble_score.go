package scrabble

import "strings"

func Score(word string) int {
  word = strings.ToUpper(word)
	count := strings.Count(word, "A")
	count += strings.Count(word, "E")
	count += strings.Count(word, "I")
	count += strings.Count(word, "O")
	count += strings.Count(word, "U")
	count += strings.Count(word, "L")
	count += strings.Count(word, "N")
	count += strings.Count(word, "R")
	count += strings.Count(word, "S")
	count += strings.Count(word, "T")
	count += strings.Count(word, "D") * 2
	count += strings.Count(word, "G") * 2

	count += strings.Count(word, "B") * 3
	count += strings.Count(word, "C") * 3
	count += strings.Count(word, "M") * 3
	count += strings.Count(word, "P") * 3

	count += strings.Count(word, "F") * 4
	count += strings.Count(word, "H") * 4
	count += strings.Count(word, "V") * 4
	count += strings.Count(word, "W") * 4
	count += strings.Count(word, "Y") * 4

	count += strings.Count(word, "K") * 5

  count += strings.Count(word, "J") * 8
	count += strings.Count(word, "X") * 8

  count += strings.Count(word, "Q") * 10
	count += strings.Count(word, "Z") * 10

  return count
}
