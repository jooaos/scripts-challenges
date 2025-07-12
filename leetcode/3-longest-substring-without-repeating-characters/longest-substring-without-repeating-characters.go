package main

func LengthOfLongestSubstring(s string) int {
	longest := 0
	inputMap := map[string]int{}
	start := 0

	for i := 0; i < len(s); i++ {
		if _, ok := inputMap[string(s[i])]; ok {
			if inputMap[string(s[i])]+1 > start {
				start = inputMap[string(s[i])] + 1
			}
		}

		inputMap[string(s[i])] = i
		if i-start+1 > longest {
			longest = i - start + 1
		}
	}

	return longest
}
