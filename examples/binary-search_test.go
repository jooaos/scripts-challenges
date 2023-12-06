package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBinarySearch(t *testing.T) {
	result := BinarySearch([]int{11, 12, 13, 14}, 13)
	assert.EqualValues(t, 2, result)
}

func TestDontFindBinarySearch(t *testing.T) {
	result := BinarySearch([]int{11, 12, 13, 14}, 15)
	assert.EqualValues(t, -1, result)
}

func TestFindWithHeavyPayloadBinarySearch(t *testing.T) {
	var i int
	var payload []int
	for i <= 100000 {
		payload = append(payload, i)
		i++
	}
	result := BinarySearch(payload, 13)
	assert.EqualValues(t, 13, result)
}
