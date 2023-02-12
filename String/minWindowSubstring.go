package string

// Leet code link -> https://leetcode.com/problems/minimum-window-substring/description/
func minWindow(s string, t string) string {

	tMap := make(map[byte]int, 0)

	tSize := len(t)
	// Fill a map with valid count of characters.
	for i := 0; i < tSize; i++ {

		tMap[t[i]]++
	}

	left := 0
	index := make([]int, 2)
	index[0] = -1
	index[1] = -1
	minWindowLen := 99999999
	sHasSize := 0

	windowMap := make(map[byte]int, 0)

	for right := 0; right < len(s); right++ {

		current := s[right]

		windowMap[current]++

		// If the current word is present in the tMap and the amount of occurance is lesser than or equal, we will update the size of the sHasSize counter.
		if checkIfInMap(tMap, current) {

			if windowMap[current] <= tMap[current] {
				sHasSize++
			}
		}

		// We have found a substring.
		if sHasSize == tSize {

			// Reset to find next window.
			for sHasSize == tSize {

				windowLen := (right - left) + 1
				//Check if it's min length and update index accordingly.
				if windowLen < minWindowLen {

					minWindowLen = windowLen
					index[0] = left
					index[1] = right
				}

				windowMap[s[left]]--

				if checkIfInMap(tMap, s[left]) {

					if windowMap[s[left]] < tMap[s[left]] {

						sHasSize--
					}
				}

				left++
			}
		}

	}

	// If index is of default value, then we don't have a valid substring.
	if index[1] == -1 || index[0] == -1 {

		return ""
	}

	// Slice and return string till the appropriate indices
	return s[index[0] : index[1]+1]
}

func checkIfInMap(hMap map[byte]int, c byte) bool {

	_, ok := hMap[c]

	return ok
}
