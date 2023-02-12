package string

import "fmt"

// Leet code link -> https://leetcode.com/problems/remove-k-digits/
func removeKdigits(num string, k int) string {

	// Handling edge cases
	if k >= len(num) {

		return "0"
	}

	// Monotonically Increasing Stack
	mIStack := IntStack{}

	idxToRemove := map[int]bool{}
	mIStack.Push(0)

	for i := 1; i < len(num); i++ {

		// Pop from stack if it is not in a monotonically increasing order
		for mIStack.Size() > 0 && rune(num[i]-'0') < rune(num[mIStack.Peek()]-'0') && k > 0 {

			idxToRemove[mIStack.Pop()] = true
			k--
		}

		// if the limit to remove is exhausted, we stop processing.
		if k == 0 {

			break
		}

		mIStack.Push(i)
	}

	finalArray := []byte{}

	for i := 0; i < len(num); i++ {

		if _, ok := idxToRemove[i]; ok {
			continue
		}

		// We do not want to append 0 in the beginning of the answer
		if len(finalArray) == 0 && rune(num[i]-'0') == 0 {

			continue
		}

		finalArray = append(finalArray, num[i])
	}

	// If k is still remaining, we have to delete from the end
	if len(finalArray) >= k {

		finalArray = finalArray[:len(finalArray)-k]
	} else {
		finalArray = []byte{}
	}

	// Handle edge case when the answer is 0
	if len(finalArray) == 0 {

		return "0"
	}

	return string(finalArray)
}

type IntStack struct {
	items []int
}

func (s *IntStack) Push(value int) {

	s.items = append(s.items, value)
}

func (s *IntStack) Pop() int {

	size := len(s.items)
	poppedItem := s.items[size-1]
	s.items = s.items[:size-1]
	return poppedItem
}

func (s *IntStack) Size() int {

	return len(s.items)
}

func (s *IntStack) Peek() int {

	return s.items[s.Size()-1]
}

func RemoveKDigitsCaller() {

	// num := "1432219"
	// k := 3

	// num := "10200"
	// k := 1

	// num := "1234567890"
	// k := 9

	// num := "12354"
	// k := 3

	num := "10001"
	k := 4

	fmt.Print(removeKdigits(num, k))
}
