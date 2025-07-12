package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LengthOfLongestSubstring(t *testing.T) {
	tc := []struct {
		input  string
		result int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
		{
			"abba",
			2,
		},
	}

	for _, teste := range tc {
		t.Run("test", func(t *testing.T) {
			result := LengthOfLongestSubstring(teste.input)
			assert.Equal(t, teste.result, result)
		})
	}
}
