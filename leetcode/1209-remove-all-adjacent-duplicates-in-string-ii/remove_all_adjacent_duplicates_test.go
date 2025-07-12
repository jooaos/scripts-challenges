package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveAllAdjacentDuplicates(t *testing.T) {
	tc := []struct {
		input  string
		times  int
		result string
	}{
		{
			"abcd",
			2,
			"abcd",
		},
		{
			"deeedbbcccbdaa",
			3,
			"aa",
		},
		{
			"pbbcggttciiippooaais",
			2,
			"ps",
		},
		{
			"abbad",
			2,
			"d",
		},
	}

	for _, test := range tc {
		t.Run("test", func(t *testing.T) {
			result := RemoveAllAdjacentDuplicates(test.input, test.times)
			assert.Equal(t, test.result, result)
		})
	}
}
