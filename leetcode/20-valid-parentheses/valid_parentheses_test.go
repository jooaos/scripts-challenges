package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidParentheses_1(t *testing.T) {
	tc := []struct {
		input  string
		output bool
	}{
		{
			input:  "()",
			output: true,
		},
		{
			input:  "()[]{}",
			output: true,
		},
		{
			input:  "(]",
			output: false,
		},
		{
			input:  "([])",
			output: true,
		},
		{
			input:  "([)]",
			output: false,
		},
	}

	for _, tst := range tc {
		t.Run("test", func(t *testing.T) {
			result := IsValid_1(tst.input)
			assert.Equal(t, tst.output, result)
		})
	}
}

func Test_ValidParentheses_2(t *testing.T) {
	tc := []struct {
		input  string
		output bool
	}{
		{
			input:  "()",
			output: true,
		},
		{
			input:  "()[]{}",
			output: true,
		},
		{
			input:  "(]",
			output: false,
		},
		{
			input:  "([])",
			output: true,
		},
		{
			input:  "([)]",
			output: false,
		},
	}

	for _, tst := range tc {
		t.Run("test", func(t *testing.T) {
			result := IsValid_2(tst.input)
			assert.Equal(t, tst.output, result)
		})
	}
}
