package main

import "sort"

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	previousNumber := 0
	maxCount := 1
	currentCount := 1
	for index, item := range nums {
		if index == 0 || item-previousNumber == 0 {
			previousNumber = item
			continue
		}

		if item-previousNumber == 1 {
			currentCount++
		} else {
			currentCount = 1
		}

		previousNumber = item

		if currentCount > maxCount {
			maxCount = currentCount
		}
	}

	return maxCount
}
