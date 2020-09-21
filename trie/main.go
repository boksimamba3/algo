package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func (tn *TrieNode) getIndex(char rune) int {
	return int(char) - int('a')
}

func (tn *TrieNode) Insert(word string) {
	current := tn
	word = strings.ToLower(word)
	for _, char := range word {
		index := tn.getIndex(char)
		if current.children[index] == nil {
			current.children[index] = &TrieNode{}
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func (tn *TrieNode) Find(word string) bool {
	current := tn
	word = strings.ToLower(word)
	for _, char := range word {
		index := tn.getIndex(char)
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}

	return current.isEnd
}

func main() {
	trie := &TrieNode{}
	trie.Insert("geek")
	trie.Insert("geeks")
	trie.Insert("bad")
	trie.Insert("bat")
	trie.Insert("batmobil")

	fmt.Println(trie.Find("bat"))
	fmt.Println(trie.Find("geek"))
	fmt.Println(trie.Find("gee"))
	fmt.Println(trie.Find("bating"))
	fmt.Println(trie.Find("batmobil"))

}
