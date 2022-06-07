package wordBreak

import (
	"fmt"
	"strings"
)

type dict map[string]bool

func newDict(words ...string) dict {
	d := dict{}
	for _, w := range words {
		d[w] = true
	}
	return d
}

func (d dict) wordBreak(s string) (broken []string, ok bool) {
	if s == "" {
		return nil, true
	}
	type prefix struct {
		length int
		broken []string
	}
	bp := []prefix{{0, nil}}
	for end := 1; end <= len(s); end++ {
		for i := len(bp) - 1; i >= 0; i-- {
			w := s[bp[i].length:end]
			if d[w] {
				b := append(bp[i].broken, w)
				if end == len(s) {
					return b, true
				}
				bp = append(bp, prefix{end, b})
				break
			}
		}
	}
	return nil, false
}

func WordBreak() {
	d := newDict("mobile", "samsung", "sam", "sung", "man", "mango",
		"icecream", "and", "go", "i", "like", "ice", "cream")
	words := []string{"ilikesamsung", "iiiiiiii", "", "samsungandmango", "ilikesamsung", "samsungandmangok"}

	for _, s := range words {
		fmt.Printf("%-20s: ", s)
		if b, ok := d.wordBreak(s); ok {
			fmt.Printf("%t: %s\n", ok, strings.Join(b, " "))
		} else {
			fmt.Println("can't break")
		}
	}
}
