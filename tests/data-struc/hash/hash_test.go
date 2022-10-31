package hash_test

import (
	"testing"
	"exercise-go/src/data-struc/hash"
	"exercise-go/tests"
)

func TestHashTable(t *testing.T) {
	table := hash.NewDataStorage()

	table.Insert("RANDOM")
	found := table.Search("RANDOM")
	tests.AssertEqual(t, found, true)
	table.Delete("RANDOM")
	found = table.Search("RANDOM")
	tests.AssertEqual(t, found, false)
}