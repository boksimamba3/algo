package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func (tn *TrieNode) getIndex(char byte) int {
	return int(char) - 'a'
}

func (tn *TrieNode) Insert(word string) {
	current := tn
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		index := tn.getIndex(word[i])
		if current.children[index] == nil {
			current.children[index] = &TrieNode{}
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func (tn *TrieNode) findNode(word string) *TrieNode {
	current := tn
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		index := tn.getIndex(word[i])
		if current.children[index] == nil {
			return nil
		}
		current = current.children[index]
	}

	return current
}

func (tn *TrieNode) Search(word string) bool {
	found := tn.findNode(word)
	if found == nil {
		return false
	}

	return found.isEnd
}

func (tn *TrieNode) isEmpty(node *TrieNode) bool {
	for i := 0; i < len(node.children); i++ {
		if node.children[i] != nil {
			return false
		}
	}

	return true
}

func (tn *TrieNode) Delete(word string) *TrieNode {
	return tn.delete(tn, word, 0)
}

func (tn *TrieNode) delete(current *TrieNode, word string, i int) *TrieNode {
	if current == nil {
		return nil
	}

	if len(word) == i {
		current.isEnd = false
		if tn.isEmpty(current) {
			current = nil
		}

		return current
	}

	index := tn.getIndex(word[i])
	current.children[index] = tn.delete(current.children[index], word, i+1)

	if tn.isEmpty(current) && !current.isEnd {
		current = nil
	}

	return current
}

func main() {
	trie := &TrieNode{}
	trie.Insert("an")
	trie.Insert("and")
	trie.Insert("ant")
	trie.Insert("bad")
	trie.Insert("bat")
	trie.Insert("zoo")

	fmt.Println("Search 'bat': ", trie.Search("bat"))
	fmt.Println("Search 'bad': ", trie.Search("bad"))
	fmt.Println("Search 'geek': ", trie.Search("geek"))
	fmt.Println("Search 'ant': ", trie.Search("ant"))
	fmt.Println("Search 'an': ", trie.Search("an"))
	fmt.Println("Search 'and': ", trie.Search("and"))
	trie.Delete("an")
	trie.Delete("and")
	fmt.Println("Search 'ant': ", trie.Search("ant"))
	fmt.Println("Search 'an': ", trie.Search("an"))
	fmt.Println("Search 'and': ", trie.Search("and"))

}
