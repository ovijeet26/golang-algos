package dynamicprogramming

import "fmt"

// The function should returna  boolean indicating whetehr or not the target can be
// contructed by concatenating elements of teh wordBank array.
// We may re-use the words as many times as needed.
func canConstruct(target string, wordBank []string) bool {

	dp := make(map[string]bool, 0)
	return canConstructMemo(target, wordBank, dp)
	//return canConstructRec(target, wordBank)
}

// Naive recursive solution
func canConstructMemo(target string, wordBank []string, dp map[string]bool) bool {

	if target == "" {
		return true
	}

	if value, ok := dp[target]; ok {

		return value
	}
	for _, word := range wordBank {

		isPrefix := true
		// Loop to check if teh current word is a valid prefix of the target.
		for i := 0; i < len(word); i++ {

			if word[i] != target[i] {
				isPrefix = false
				break
			}
		}

		if isPrefix {

			// Slice if off from the word's last index
			remainingSuffix := target[len(word):]

			// If the suffix can be constructed,a nd we know the taget can be constructed using teh current word,
			// we would forcibly return tru to end the cycle.
			if canConstructMemo(remainingSuffix, wordBank, dp) {
				dp[remainingSuffix] = true
				return true
			}
		}
	}

	dp[target] = false
	return false
}

// Naive recursive solution
func canConstructRec(target string, wordBank []string) bool {

	if target == "" {
		return true
	}

	for _, word := range wordBank {

		isPrefix := true
		// Loop to check if teh current word is a valid prefix of the target.
		for i := 0; i < len(word); i++ {

			if word[i] != target[i] {
				isPrefix = false
				break
			}
		}

		if isPrefix {

			// Slice if off from the word's last index
			remainingSuffix := target[len(word):]

			// If the suffix can be constructed, and we know the taget can be constructed using the current word,
			// we would forcibly return true to end the cycle.
			if canConstruct(remainingSuffix, wordBank) {
				return true
			}
		}
	}

	return false
}

func CanConstructCaller() {

	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeeee", "eeeeeee"}))

}
