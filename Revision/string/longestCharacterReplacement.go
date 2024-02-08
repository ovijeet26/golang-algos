package string

import "fmt"

func characterReplacement(s string, k int) int {

	size := len(s)
	if size <= 1 {
		return size
	}

	maxWindow := 1
	charCount := make(map[rune]int, 0)
	left, right := 0, 0
	charCount[rune(s[left])]++
	var maxCount int

	for right < size {
		maxCount = 0

		for _, value := range charCount {

			maxCount = maxValue(maxCount, value)
		}

		windowSize := right - left + 1

		if windowSize-maxCount <= k {
			maxWindow = maxValue(maxWindow, windowSize)
			right++
			if right >= size {
				break
			}
			charCount[rune(s[right])]++
			continue
		}

		// Window size is not valid
		charCount[rune(s[left])]--
		left++
	}

	return maxWindow
}

func maxValue(a int, b int) int {

	if a > b {
		return a
	}

	return b
}

func LongestCharRCaller() {

	fmt.Println(characterReplacement("AABABBA", 1))
}
