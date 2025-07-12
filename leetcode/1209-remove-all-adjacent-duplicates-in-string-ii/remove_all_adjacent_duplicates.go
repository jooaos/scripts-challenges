package main

func RemoveAllAdjacentDuplicates(s string, k int) string {
	chars := []byte{}
	counts := []int{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if len(chars) == 0 || chars[len(chars)-1] != c {
			chars = append(chars, c)
			counts = append(counts, 1)
		} else {
			counts[len(counts)-1]++
			if counts[len(counts)-1] == k {
				chars = chars[:len(chars)-1]
				counts = counts[:len(counts)-1]
			}
		}
	}

	result := make([]byte, 0)
	for i := 0; i < len(chars); i++ {
		for j := 0; j < counts[i]; j++ {
			result = append(result, chars[i])
		}
	}

	return string(result)
}
