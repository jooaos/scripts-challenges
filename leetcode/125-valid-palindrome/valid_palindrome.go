package main

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	s = sanitaizeString(s)

	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func sanitaizeString(s string) string {
	s = strings.ToLower(s)
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	return reg.ReplaceAllString(s, "")
}
