package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestConsecutive(t *testing.T) {
	tc := []struct {
		input  *ListNode
		result *ListNode
	}{
		{
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
			&ListNode{
				5,
				&ListNode{
					4,
					&ListNode{
						3,
						&ListNode{
							2,
							&ListNode{
								1,
								nil,
							},
						},
					},
				},
			},
		},
	}

	for _, teste := range tc {
		t.Run("test", func(t *testing.T) {
			result := ReverseList(teste.input)

			i := 5
			for result != nil {
				assert.Equal(t, i, result.Val)
				result = result.Next
				i--
			}
		})
	}
}
