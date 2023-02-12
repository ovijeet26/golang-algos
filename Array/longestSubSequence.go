package array

// Leet code link -> https://leetcode.com/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {

	// Implement a set in golang.
	set := make(map[int]struct{})
	present := struct{}{}

	seqLength := 1
	maxSeqLength := 0

	size := len(nums)

	// Add all the nums in the set.
	for i := 0; i < size; i++ {

		set[nums[i]] = present
	}

	// Iterate nums and check if the current item has a left neighbour.
	for i := 0; i < size; i++ {

		// If the current item has a left neighbour, then it is NOT the beginning of a sub-sequence.
		if isNumberPresentInSet(set, nums[i]-1) {

			continue
		}

		currentNumber := nums[i]

		// If it doesn't then it's the start of a sequence.
		// Using the set, check the length of the sequence. Note the max count.
		for isNumberPresentInSet(set, currentNumber+1) {

			seqLength++
			currentNumber++
		}
		maxSeqLength = max(maxSeqLength, seqLength)

		// Reset the sequence length
		seqLength = 1
	}

	return maxSeqLength
}

// Check if the number is present inside the set
func isNumberPresentInSet(set map[int]struct{}, value int) bool {

	_, ok := set[value]

	return ok
}

// func max(a int, b int) int {

// 	if a > b {

// 		return a
// 	}

// 	return b
// }
