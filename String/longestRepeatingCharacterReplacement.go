package string

import "fmt"

// Leet code link -> https://leetcode.com/problems/longest-repeating-character-replacement/
// Solution reference -> https://www.youtube.com/watch?v=gqXU1UyA8pk&ab_channel=NeetCode
func characterReplacement(s string, k int) int {

	left := 0

	windowMap := map[byte]int{}
	result := 0

	for right := 0; right < len(s); right++ {

		// Add the element frequency inside the windowMap
		windowMap[s[right]]++

		// Calculate the window length
		windowLength := (right - left) + 1

		maxCharFrequency := maxFrequency(windowMap)

		// If the amount of characters to be replaced have to be greater than k then we shrink the window
		if windowLength-maxCharFrequency > k {

			windowMap[s[left]]--
			left++
		}

		validWindowLength := (right - left) + 1
		if validWindowLength > result {

			result = validWindowLength
		}
	}

	return result
}

// Calculate max freuqency of a word within the map
func maxFrequency(hashMap map[byte]int) int {

	max := 0
	for _, count := range hashMap {

		if count > max {
			max = count
		}
	}
	return max
}

func CharacterReplacementCaller() {

	num := characterReplacement("ABAA", 0)

	fmt.Print(num)
}
