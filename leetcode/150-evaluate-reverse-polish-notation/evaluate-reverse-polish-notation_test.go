package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EvalRPN(t *testing.T) {
	tc := []struct {
		input  []string
		result int
	}{
		{
			[]string{"2", "1", "+", "3", "*"},
			9,
		},
		{
			[]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"},
			22,
		},
	}

	for _, teste := range tc {
		t.Run("test", func(t *testing.T) {
			result := EvalRPN(teste.input)
			assert.Equal(t, teste.result, result)
		})
	}
}
