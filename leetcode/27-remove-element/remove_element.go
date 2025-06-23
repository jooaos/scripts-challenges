package main

func RemoveElement_1(nums []int, val int) int {
	numbersToOutput := []int{}
	for _, item := range nums {
		if item != val {
			numbersToOutput = append(numbersToOutput, item)
		}
	}

	for index, item := range numbersToOutput {
		nums[index] = item
	}

	return len(numbersToOutput)
}

func RemoveElement_2(nums []int, val int) int {
	index := 0
	for _, item := range nums {
		if item != val {
			nums[index] = item
			index++
		}
	}

	return index
}
