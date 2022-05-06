package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	result := make([]string, 0)
	subject = strings.ToLower(subject)
	sc := strings.Split(subject, "")
	sort.Strings(sc)
	for _, candidate := range candidates {
		candidateLC := strings.ToLower(candidate)
		cc := strings.Split(candidateLC, "")
		sort.Strings(cc)
		keep := true
		if len(sc) != len(cc) || candidateLC == subject {
			keep = false
		} else {
			for idx := range sc {
				if sc[idx] != cc[idx] {
					keep = false
					break
				}
			}
		}
		if keep {
			result = append(result, candidate)
		}
	}
	return result
}
