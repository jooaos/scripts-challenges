package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestConsecutive(t *testing.T) {
	tc := []struct {
		input_1 []int
		input_2 int
		result  []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			2,
			[]int{1, 2},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, teste := range tc {
		t.Run("test", func(t *testing.T) {
			result := TopKFrequent(teste.input_1, teste.input_2)
			assert.Equal(t, teste.result, result)
		})
	}
}
