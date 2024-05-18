package string

func longestPalindrome(s string) string {

	res := ""
	resLen := 0

	for i := 0; i <= len(s); i++ {

		// Check for cases for odd length palindromes e.g. bab
		left, right := i, i

		for left >= 0 && right < len(s) && s[left] == s[right] {

			lenOfPalindrome := right - left + 1
			if resLen < lenOfPalindrome {
				resLen = lenOfPalindrome
				res = s[left : right+1]
			}
			left--
			right++
		}

		// Check for cases for odd length palindromes e.g. bab
		left, right = i, i+1

		for left >= 0 && right < len(s) && s[left] == s[right] {

			lenOfPalindrome := right - left + 1
			if resLen < lenOfPalindrome {
				resLen = lenOfPalindrome
				res = s[left : right+1]
			}
			left--
			right++
		}
	}

	return res
}
