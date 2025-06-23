package main

func RemoveDuplicates(nums []int) int {
	uniqueElements := []int{}
	var previousElement int

	for index, item := range nums {
		if item != previousElement || index == 0 {
			uniqueElements = append(uniqueElements, item)
		}

		previousElement = item
	}

	for i := 0; i < len(uniqueElements); i++ {
		nums[i] = uniqueElements[i]
	}

	return len(uniqueElements)
}
