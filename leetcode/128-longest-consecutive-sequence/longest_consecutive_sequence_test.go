package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestConsecutive(t *testing.T) {
	tc := []struct {
		input  []int
		result int
	}{
		{
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
		{
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			9,
		},
		{
			[]int{1, 0, 1, 2},
			3,
		},
	}

	for _, teste := range tc {
		t.Run("test", func(t *testing.T) {
			result := LongestConsecutive(teste.input)
			assert.Equal(t, teste.result, result)
		})
	}
}
