package main

func GroupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)

	for i := 0; i < len(strs); i++ {
		var alphabet [26]int
		for _, letter := range strs[i] {
			alphabet[letter-'a']++
		}
		anagramMap[alphabet] = append(anagramMap[alphabet], strs[i])
	}

	output := [][]string{}

	for _, item := range anagramMap {
		output = append(output, item)
	}

	return output
}
