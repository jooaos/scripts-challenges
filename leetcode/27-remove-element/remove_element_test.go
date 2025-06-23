package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	tc := []struct {
		inputArray []int
		inputVal   int
		result     int
	}{
		{
			inputArray: []int{3, 2, 2, 3},
			inputVal:   3,
			result:     2,
		},
		{
			inputArray: []int{0, 1, 2, 2, 3, 0, 4, 2},
			inputVal:   2,
			result:     5,
		},
	}

	for _, test := range tc {
		t.Run("teste", func(t *testing.T) {
			result := RemoveElement_1(test.inputArray, test.inputVal)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestRemoveElement2(t *testing.T) {
	tc := []struct {
		inputArray []int
		inputVal   int
		result     int
	}{
		{
			inputArray: []int{3, 2, 2, 3},
			inputVal:   3,
			result:     2,
		},
		{
			inputArray: []int{0, 1, 2, 2, 3, 0, 4, 2},
			inputVal:   2,
			result:     5,
		},
	}

	for _, test2 := range tc {
		t.Run("teste", func(t *testing.T) {
			result := RemoveElement_2(test2.inputArray, test2.inputVal)
			assert.Equal(t, test2.result, result)
		})
	}
}
