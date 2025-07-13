package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestConsecutive(t *testing.T) {
	tc := []struct {
		input   *TreeNode
		input_2 int
		result  *TreeNode
	}{
		{
			&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			2,
			&TreeNode{},
		},
		{
			&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			10,
			nil,
		},
	}

	for _, teste := range tc {
		t.Run("test", func(t *testing.T) {
			result := SearchBST(teste.input, teste.input_2)
			if teste.result != nil {
				assert.NotNil(t, result)
			} else {
				fmt.Println(result)
				assert.Nil(t, result)
			}
		})
	}
}
