package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessRecursiveSum(t *testing.T) {
	result := RecursiveSum([]int{1, 2, 3})
	assert.EqualValues(t, 6, result)
}
