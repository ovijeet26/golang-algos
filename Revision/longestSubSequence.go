package revision

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	longestSequence, localSequence := 1, 0
	hashMap := make(map[int]bool, 0)

	for _, v := range nums {
		hashMap[v] = true
	}

	for _, num := range nums {

		localSequence = 0

		// is the current number part of a sequence, then ignore
		_, prevPresent := hashMap[num-1]

		if prevPresent {
			continue
		}

		// is the current number start of a sequence, then count the total sequence length
		_, nextPresent := hashMap[num+1]

		if nextPresent {
			localSequence++
			ptr := num + 1
			_, nextPresent = hashMap[ptr]
			for nextPresent {

				localSequence++
				ptr++
				_, nextPresent = hashMap[ptr]
			}
		}

		longestSequence = maxNum(longestSequence, localSequence)
	}
	return longestSequence
}

func maxNum(a int, b int) int {

	if a < b {
		return b
	}

	return a
}
