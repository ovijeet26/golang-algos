package string

func countSubstrings(s string) int {

	resCount := 0

	for i := 0; i <= len(s); i++ {

		left, right := i, i+1

		for left >= 0 && right < len(s) && s[left] == s[right] {
			resCount++
			right++
			left--
		}

		left, right = i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			resCount++
			right++
			left--
		}
	}

	return resCount
}
