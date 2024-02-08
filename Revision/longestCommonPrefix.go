package revision

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

func longestCommonPrefix2(strs []string) string {
	result := ""
	for _, str := range strs {
		if str == "" {
			return ""
		} else if result == "" {
			result += str
		} else {
			if len(str) < len(result) {
				result = result[:len(str)]
			}
			for pos := range str {
				if pos == len(result) {
					break
				} else if str[pos] != result[pos] {
					result = result[:pos]
					break
				}
			}

			if result == "" {
				return ""
			}
		}
	}
	return result
}
