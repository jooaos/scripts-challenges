package main

func DailyTemperatures(temperatures []int) []int {
	output := make([]int, len(temperatures))
	stack := []int{}
    
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			output[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}

	return output
}