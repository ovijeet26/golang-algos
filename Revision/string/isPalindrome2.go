package string

func validPalindrome(s string) bool {

	size := len(s)

	if size <= 1 {
		return true
	}

	left, right := 0, size-1
	for left < right {

		if s[left] == s[right] {
			left++
			right--
			continue
		}

		isLeftPalindrome, isRightPalindrome := false, false
		if s[left+1] == s[right] {
			isLeftPalindrome = ispalindrome(s[left+1 : right+1])
		}

		if !isLeftPalindrome && s[left] == s[right-1] {
			isRightPalindrome = ispalindrome(s[left:right])
		}

		if !isLeftPalindrome && !isRightPalindrome {
			return false
		}
		return true
	}
	return true
}

func ispalindrome(s string) bool {

	size := len(s)

	if size <= 1 {
		return true
	}

	left, right := 0, size-1

	for left < right {

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}
	return true
}
