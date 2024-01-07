package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	result := QuickSort([]int{14, 12, 13, 11})
	assert.EqualValues(t, []int{11, 12, 13, 14}, result)
}

func TestQuickSortWithValuesDuplicated(t *testing.T) {
	result := QuickSort([]int{3, 7, 1, 4, 6, 1, 2, 8, 8, 5, 9})
	assert.EqualValues(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
}
