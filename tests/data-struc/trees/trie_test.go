package trees_test

import (
	"exercise-go/src/data-struc/trees"
	"exercise-go/tests"
	"testing"
)

func TestTires(t *testing.T) {
	trie := trees.NewTrie()
	trie.Insert("Manga")
	trie.Insert("Mango")

	tests.AssertEqual(t, true, trie.Search("Manga"))
	tests.AssertEqual(t, false, trie.Search("Mang"))
	tests.AssertEqual(t, false, trie.Search("Meng"))
	tests.AssertEqual(t, false, trie.Search("Mangal"))
	tests.AssertEqual(t, true, trie.Search("Mango"))
}
