package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestConsecutive(t *testing.T) {
	tc := []struct {
		input  []int
		result []int
	}{
		{
			[]int{73, 74, 75, 71, 69, 72, 76, 73},
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			[]int{30, 40, 50, 60},
			[]int{1, 1, 1, 0},
		},
	}

	for _, teste := range tc {
		t.Run("test", func(t *testing.T) {
			result := DailyTemperatures(teste.input)
			assert.Equal(t, teste.result, result)
		})
	}
}
