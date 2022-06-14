package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}

	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	for i := 0; i < len(wordsToFind); i++ {
		found := trie.Search(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}
	fmt.Printf("Search Prefix \"%s\" found %t\n", "jos",trie.StartsWith("jos"))
}
