package main

func TopKFrequent(nums []int, k int) []int {
	mapFreq := map[int]int{}
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		mapFreq[num]++
	}

	for num, f := range mapFreq {
		freq[f] = append(freq[f], num)
	}

	result := []int{}
	for i := len(freq) - 1; i >= 0; i-- {
		for _, num := range freq[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}
