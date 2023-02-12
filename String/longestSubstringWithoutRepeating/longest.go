package longestsubstringwithoutrepeating

//Given a string s, find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {

	// Solution using sliding window technique.
	longestGlobalCount := 0
	longestLocalCount := 1

	indexHashMap := make(map[byte]int, 0)
	left := 0
	right := 0

	for right <= len(s)-1 {

		index, ok := indexHashMap[s[right]]

		if ok && index >= left {
			left = index + 1
			longestLocalCount = 0
		} else {

			indexHashMap[s[right]] = right

			longestLocalCount = (right - left) + 1
			right++

			if longestLocalCount > longestGlobalCount {
				longestGlobalCount = longestLocalCount
			}
		}
	}

	return longestGlobalCount
}