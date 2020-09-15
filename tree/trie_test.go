package tree

import (
	"fmt"
	"testing"
)

var (
	trie = NewTrie()
)

func TestTrie_Insert(t *testing.T) {
	trie.Insert("apple")
	fmt.Printf("%v\n", trie.Search("apple"))
	fmt.Printf("%v\n", trie.Search("app"))
	fmt.Printf("%v\n", trie.StartWith("app"))
	trie.Insert("app")
	fmt.Printf("%v\n", trie.Search("app"))
}
