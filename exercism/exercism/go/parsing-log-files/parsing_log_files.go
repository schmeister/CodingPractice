package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	r, _ := regexp.Compile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
	val := r.MatchString(text)
	return val
}

func SplitLogLine(text string) []string {
	r, _ := regexp.Compile(`<[-=~*]*>`)
	strings := r.Split(text, -1)
	return strings
}

func CountQuotedPasswords(lines []string) int {
	passReg, _ := regexp.Compile(`(?i)".*password.*"`)
	count := 0
	for _, t := range lines {
		if passReg.MatchString(t) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	eolReg, _ := regexp.Compile(`end-of-line[0-9]*`)
	stripped := eolReg.ReplaceAllString(text, "")
	return stripped
}

func TagWithUserName(lines []string) []string {
	userReg, _ := regexp.Compile(`User [ ]*`)
	nameReg, _ := regexp.Compile(`[A-Z][a-z]+[0-9]+`)
	//	nameReg, _ := regexp.Compile(`(?<=\bUser\s)(\w+)`)
	for idx, line := range lines {
		if userReg.MatchString(line) {
			name := nameReg.FindString(line)
			new := fmt.Sprintf("[USR] %s %s", name, line)
			lines[idx] = new
		}
	}
	return lines
}
