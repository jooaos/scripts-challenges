package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canSchedule([][]int{
		{
			0, 30,
		},
		{
			5, 10,
		},
		{
			15, 20,
		},
	}))

	fmt.Println(canSchedule([][]int{
		{
			7, 10,
		},
		{
			2, 4,
		},
	}))
}

func canSchedule(input [][]int) bool {
	if len(input) == 0 {
		return true
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i][0] < input[j][0]
	})

	for i := 1; i < len(input); i++ {
		if input[i][0] < input[i-1][1] {
			return false
		}
	}

	return true
}
