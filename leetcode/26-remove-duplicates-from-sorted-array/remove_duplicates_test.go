package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tc := []struct {
		input  []int
		result int
	}{
		{
			input:  []int{1, 1, 2},
			result: 2,
		},
		{
			input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			result: 5,
		},
	}

	for _, test := range tc {
		t.Run("test", func(t *testing.T) {
			r := RemoveDuplicates(test.input)
			assert.Equal(t, test.result, r)
		})
	}
}
