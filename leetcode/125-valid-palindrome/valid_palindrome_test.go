package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestConsecutive(t *testing.T) {
	tc := []struct {
		input  string
		result bool
	}{
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
		{
			"",
			true,
		},
	}

	for _, teste := range tc {
		t.Run("test", func(t *testing.T) {
			result := IsPalindrome(teste.input)
			assert.Equal(t, teste.result, result)
		})
	}
}
