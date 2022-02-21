package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedMapKeysSortsKeys(t *testing.T) {
	mapOfInts := map[string]int{
		"A": 1,
		"Z": 1,
		"B": 1,
	}

	expected := []string{
		"A", "B", "Z",
	}

	assert.Equal(t, expected, sortedMapKeys(mapOfInts))
}
