package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestConsecutive(t *testing.T) {
	tc := []struct {
		input_1 string
		input_2 string
		result  bool
	}{
		{
			"anagram",
			"nagaram",
			true,
		},
		{
			"rat",
			"car",
			false,
		},
	}

	for _, teste := range tc {
		t.Run("test", func(t *testing.T) {
			result := IsAnagram(teste.input_1, teste.input_2)
			assert.Equal(t, teste.result, result)
		})
	}
}
