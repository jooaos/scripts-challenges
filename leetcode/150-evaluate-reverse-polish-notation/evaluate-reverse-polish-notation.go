package main

import (
	"fmt"
	"strconv"
)

func EvalRPN(tokens []string) int {
	stackOperations := []string{}
	result := 0

	for i := 0; i < len(tokens); i++ {
		if i == 0 {
			result, _ = strconv.Atoi(tokens[0])
		}
		if isOperator(tokens[i]) {
			result = calculate(stackOperations[len(stackOperations)-1], stackOperations[len(stackOperations)-2], tokens[i])
			stackOperations = stackOperations[:len(stackOperations)-2]
			stackOperations = append(stackOperations, fmt.Sprintf("%d", result))
		} else {
			stackOperations = append(stackOperations, tokens[i])
		}
	}

	return result
}

func isOperator(operator string) bool {
	switch operator {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	default:
		return false
	}
}

func calculate(num1Str, num2Str, op string) int {
	num1, _ := strconv.Atoi(num1Str)

	num2, _ := strconv.Atoi(num2Str)

	switch op {
	case "+":
		return num1 + num2
	case "-":
		return num2 - num1
	case "*":
		return num1 * num2
	case "/":
		return num2 / num1
	default:
		return 0
	}
}
