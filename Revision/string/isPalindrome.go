package string

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {

	size := len(s)

	if size <= 1 {
		return true
	}

	s = strings.ToUpper(s)

	left, right := 0, size-1

	for left < right {

		for left < right && (!unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left]))) {
			left++
		}

		for left < right && (!unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right]))) {
			right--
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
