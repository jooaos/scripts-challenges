package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tc := []struct {
		input  []string
		output [][]string
	}{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			output: [][]string{
				{
					"eat",
					"tea",
					"ate",
				},
				{
					"tan",
					"nat",
				},
				{
					"bat",
				},
			},
		},
	}

	for _, test := range tc {
		t.Run("test", func(t *testing.T) {
			result := GroupAnagrams(test.input)
			assert.Equal(t, test.output, result)
		})
	}
}
