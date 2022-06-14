package trie

type TrieNode struct {
	Children [26]*TrieNode
	IsWord   bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() *Trie {
	trie := &Trie{
		root: &TrieNode{},
	}
	return trie
}

func (trie *Trie) Insert(word string) {
	wordLength := len(word)
	current := trie.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.Children[index] == nil {
			current.Children[index] = &TrieNode{}
		}
		current = current.Children[index]
	}
	current.IsWord = true
}

func (trie *Trie) Search(word string) bool {
	wordLength := len(word)
	current := trie.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.Children[index] == nil {
			return false
		}
		current = current.Children[index]
	}
	return current.IsWord
}

func (trie *Trie) StartsWith(prefix string) bool {
	wordLength := len(prefix)
	current := trie.root
	i := 0
	for i = 0; i < wordLength; i++ {
		index := prefix[i] - 'a'
		if current.Children[index] == nil {
			return false
		}
		current = current.Children[index]
	}
	return i == wordLength
}
