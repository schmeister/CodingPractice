package ocr

import (
	"strings"
)

const underscore = "_"
const pipe = "|"
const space = " "

// Recognize returns a slice of strings containing the lines of numbers detected
// in the provided input.
func Recognize(s string) []string {
	var o []string
	lines := strings.Split(s, "\n")
	numLines := len(lines)

	// Each group is 4 lines in height, so skip the next 4 to process.
	for i := 1; i < numLines; i += 4 {
		//		fmt.Printf("found: %s\n", recognizeLine(lines[i:i+3]))
		o = append(o, recognizeLine(lines[i:i+3]))
	}
	return o
}
func recognizeLine(line []string) string {
	var o string
	for i := 0; i < len(line[0]); i += 3 {
		lines := []string{line[0][i : i+3], line[1][i : i+3], line[2][i : i+3]}
		o = o + recognizeDigit(lines)
	}
	return o
}

// recognizeDigit - for fun, gonna recognize individual items, not whole sections.
// Not sure if this is faster, but more challenging.
func recognizeDigit(line []string) string {
	first := strings.Split(line[0][0:3], "")
	second := strings.Split(line[1][0:3], "")
	third := strings.Split(line[2][0:3], "")

	// 1 or 4 - row(0) = space
	if first[1] == space {
		if second[0] == space && second[1] == space && second[2] == pipe {
			return "1"
		} else if second[0] == pipe && second[1] == underscore && second[2] == pipe {
			return "4"
		} else {
			return "?"
		}
	} else {
		// 2, 3 - space underscore space
		if second[0] == space && second[1] == underscore && second[2] == pipe {
			if third[0] == pipe {
				return "2"
			} else {
				return "3"
			}
		} else if second[0] == pipe && second[1] == underscore && second[2] == space {
			// 5, 6 - pipe underscore space
			if third[0] == space && third[1] == underscore && third[2] == pipe {
				return "5"
			} else {
				return "6"
			}
		} else if second[0] == space && second[1] == space && second[2] == pipe {
			return "7"
			// 7 - space space pipe
		} else if second[0] == pipe && second[1] == underscore && second[2] == pipe {
			// 8, 9 - pipe underscore pipe
			if third[0] == pipe && third[1] == underscore && third[2] == pipe {
				return "8"
			} else {
				return "9"
			}
		} else if second[0] == pipe && second[1] == space && second[2] == pipe {
			// 0
			if third[0] == pipe && third[1] == underscore && third[2] == pipe {
				return "0"
			} else {
				return "?"
			}
			
		}

	}
	return "?"
}
