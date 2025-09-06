package main

import (
	"sort"
)

func TotalDays(input [][]int) int {
	if len(input) == 0 {
		return 0
	}

	line := map[int]int{}

	for _, period := range input {
		line[period[0]]++
		line[period[1]]--
	}

	keys := make([]int, 0, len(line))
	for key := range line {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	total := 0
	max := 1
	for _, key := range keys {
		total += line[key]
		if total > max {
			max = total
		}
	}

	return max
}
