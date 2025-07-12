package main

func IsAnagram(s string, t string) bool {
	sMap := make(map[rune]int)

	for _, ch := range s {
		sMap[ch]++
	}

	for _, ch := range t {
		if _, ok := sMap[ch]; !ok {
			return false
		}

		sMap[ch]--

		if sMap[ch] == 0 {
			delete(sMap, ch)
		}
	}

	if len(sMap) > 0 {
		return false
	}

	return true
}
