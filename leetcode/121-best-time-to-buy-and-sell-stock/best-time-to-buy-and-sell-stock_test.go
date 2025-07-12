package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LengthOfLongestSubstring(t *testing.T) {
	tc := []struct {
		input  []int
		result int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, teste := range tc {
		t.Run("test", func(t *testing.T) {
			result := MaxProfit(teste.input)
			assert.Equal(t, teste.result, result)
		})
	}
}
