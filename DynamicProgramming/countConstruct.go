package dynamicprogramming

import "fmt"

// the function should return the number of ways that the 'target' ca
// be constructed by concatenating elements of the word bank array.
// We may reuse the words of wod bank as many times as needed.

func countConstruct(target string, wordBank []string) int {

	dp := make(map[string]int, 0)
	count := countConstructMemo(target, wordBank, dp)
	// count := countConstructRec(target, wordBank)
	return count
}

// Memoised recursive solution
func countConstructMemo(target string, wordBank []string, dp map[string]int) int {

	if target == "" {

		return 1
	}

	if value, ok := dp[target]; ok {

		return value
	}

	total := 0

	for _, word := range wordBank {

		isPrefix := true
		// Loop to check if the current word is a valid prefix of the target.
		for i := 0; i < len(word); i++ {

			if word[i] != target[i] {
				isPrefix = false
				break
			}
		}

		if isPrefix {

			// Slice if off from the word's last index
			remainingSuffix := target[len(word):]

			numWaysForRest := countConstructMemo(remainingSuffix, wordBank, dp)

			dp[remainingSuffix] = numWaysForRest

			total += numWaysForRest
		}
	}

	dp[target] = total
	return total
}

// Naive recursive solution
func countConstructRec(target string, wordBank []string) int {

	if target == "" {
		return 1
	}

	total := 0
	for _, word := range wordBank {

		isPrefix := true
		// Loop to check if the current word is a valid prefix of the target.
		for i := 0; i < len(word); i++ {

			if word[i] != target[i] {
				isPrefix = false
				break
			}
		}

		if isPrefix {

			// Slice if off from the word's last index
			remainingSuffix := target[len(word):]

			numWaysForRest := countConstructRec(remainingSuffix, wordBank)

			total += numWaysForRest
		}
	}

	return total
}

func CountConstructCaller() {

	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef"}))

	fmt.Println(countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeeee", "eeeeeee"}))

}
