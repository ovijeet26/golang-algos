package string

// Leet code link -> https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {

	sSize := len(s)
	tSize := len(t)

	if sSize != tSize {

		return false
	}

	charMap := make(map[byte]int, 0)

	for i := 0; i < sSize; i++ {

		charMap[s[i]]++
	}

	for i := 0; i < tSize; i++ {

		if _, ok := charMap[t[i]]; !ok {

			return false
		}

		charMap[t[i]]--

		if charMap[t[i]] == 0 {

			delete(charMap, t[i])
		}
	}

	// if the length of charMap still has some values, then the string is not an anagram.
	return len(charMap) == 0
}
