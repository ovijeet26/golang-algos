package dynamicprogramming

// // The function should return a 2d array containing  all of the ways that the 'target'
// // can be constructed by concatenating elements of the 'wordBank' array.
// // We may re-use elements of 'wordbank' as many times as needed.
// func allConstruct(target string, wordBank []string) [][]string {

// }

// func allConstructRec(target string, wordBank []string) [][]string {

// 	if target == "" {

// 		return [][]string{}
// 	}

// 	result := make([][]string, 0)

// 	for _, word := range wordBank {

// 		isPrefix := true
// 		// Loop to check if teh current word is a valid prefix of the target.
// 		for i := 0; i < len(word); i++ {

// 			if word[i] != target[i] {
// 				isPrefix = false
// 				break
// 			}
// 		}

// 		if isPrefix {

// 			// Slice if off from the word's last index
// 			remainingSuffix := target[len(word):]

// 			resultSuffixArray := allConstruct(remainingSuffix, wordBank)
// 			currentArray := []string{word}
// 		}
// 	}
// }
