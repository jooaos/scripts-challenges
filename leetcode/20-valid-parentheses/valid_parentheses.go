package main

func IsValid_1(s string) bool {
	brackets1 := []string{"(", ")"}
	brackets2 := []string{"{", "}"}
	brackets3 := []string{"[", "]"}

	openBrackets1, openBrackets2, openBrackets3 := 0, 0, 0
	openBrackets := []string{}
	idealBracket := ""

	for i := 0; i < len(s); i++ {
		currentBracket := string(s[i])
		isRemoved := false

		switch currentBracket {
		case brackets1[0]:
			openBrackets1++
			openBrackets = append(openBrackets, brackets1[0])
		case brackets1[1]:
			openBrackets1--
			idealBracket = brackets1[0]
			isRemoved = true
		case brackets2[0]:
			openBrackets2++
			openBrackets = append(openBrackets, brackets2[0])
		case brackets2[1]:
			openBrackets2--
			idealBracket = brackets2[0]
			isRemoved = true
		case brackets3[0]:
			openBrackets3++
			openBrackets = append(openBrackets, brackets3[0])
		case brackets3[1]:
			openBrackets3--
			idealBracket = brackets3[0]
			isRemoved = true
		}

		if isRemoved {
			if len(openBrackets) == 0 {
				return false
			}

			lastOpenBracket := openBrackets[len(openBrackets)-1]
			if lastOpenBracket != idealBracket {
				return false
			}

			openBrackets = openBrackets[:len(openBrackets)-1]
		}

		if openBrackets1 < 0 || openBrackets2 < 0 || openBrackets3 < 0 {
			return false
		}
	}

	return openBrackets1 == 0 && openBrackets2 == 0 && openBrackets3 == 0
}

func IsValid_2(s string) bool {
	bracketsMapping := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	brackets := []string{}

	for i := 0; i < len(s); i++ {
		currentBracket := string(s[i])

		if _, ok := bracketsMapping[currentBracket]; ok {
			if len(brackets) == 0 {
				return false
			}
			lastOpenedBracket := brackets[len(brackets)-1]
			if lastOpenedBracket != bracketsMapping[currentBracket] {
				return false
			}
			brackets = brackets[:len(brackets)-1]
		} else {
			brackets = append(brackets, currentBracket)
		}
	}

	return len(brackets) == 0
}
