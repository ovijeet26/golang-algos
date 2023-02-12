package almostpalindrome

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {

	s = sanitizeString(s)

	stringLength := len(s)

	left, right := 0, stringLength-1

	for left < right {

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func sanitizeString(s string) string {

	stringSanitizer := regexp.MustCompile("[^A-Za-z0-9]+")

	s = stringSanitizer.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	return s
}
