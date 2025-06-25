package main

func ProductExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	pos := make([]int, len(nums))

	pre[0] = nums[0]
	pos[len(nums)-1] = nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[i]
	}

	for i := len(nums) - 2; i > 0; i-- {
		pos[i] = pos[i+1] * nums[i]
	}

	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		previousPosition := i - 1
		previousResult := 1
		if previousPosition >= 0 {
			previousResult = pre[previousPosition]
		}

		posPosition := i + 1
		posResult := 1
		if posPosition < len(nums) {
			posResult = pos[posPosition]
		}

		result[i] = previousResult * posResult
	}

	return result
}
