package almostpalindrome

import (
	"regexp"
	"strings"
)

// Given a string s, return true if the s can be palindrome after deleting at most one character from it.

func ValidPalindrome(s string) bool {

	stringLength := len(s)

	s = sanitizeStringx(s)
	left, right := 0, stringLength-1

	// Taking case of edge cases
	if right-left < 2 {

		return true
	}

	for left < right {

		if s[left] != s[right] {

			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
		left++
		right--
	}
	return true
}

func isPalindrome(str string, left int, right int) bool {

	for left < right {

		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func sanitizeStringx(str string) string {

	stringSanitizer := regexp.MustCompile("[^a-zA-Z0-9]+")

	str = stringSanitizer.ReplaceAllString(str, "")

	str = strings.ToLower(str)

	return str
}
