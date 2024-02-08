package array

func longestCommonPrefix(strs []string) string {

	minLen := len(strs[0])

	prefix := ""
	// Calculate the minimum length of the array of strings
	for _, str := range strs {

		if len(str) < minLen {
			minLen = len(str)
		}
	}

	isMismatched := false
	for i := 0; i < minLen; i++ {

		currentChar := strs[0][i]
		for _, str := range strs {
			if str[i] != currentChar {
				isMismatched = true
				break
			}
		}

		if isMismatched {
			break
		}

		prefix = strs[0][:i+1]
	}

	return prefix
}
