package stack

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {

	stack := Stack{}

	var currentNum int
	var currentStr strings.Builder
	i := 0
	for i < len(s) {
		currentStr.Reset()
		currentNum = 0
		// add all the digits in the proper format and add to stack
		if unicode.IsDigit(rune(s[i])) {

			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				digit, _ := strconv.Atoi(string(s[i]))
				currentNum = currentNum*10 + digit
				i++
			}
			stack.Push(fmt.Sprintf("%d", currentNum))
			continue
		}

		// add all the letters in the proper format and add to stack
		if unicode.IsLetter(rune(s[i])) {

			for i < len(s) && unicode.IsLetter(rune(s[i])) {
				currentStr.WriteRune(rune(s[i]))
				i++
			}
			stack.Push(currentStr.String())
			continue
		}

		if rune(s[i]) == '[' {

			stack.Push(string(rune(s[i])))
			i++
			continue
		}

		if rune(s[i]) == ']' {

			var str string
			// Find the string to make the final string
			for stack.Length() > 0 && isCombinationOfLetters(stack.Peek()) {
				str = stack.Pop() + str
			}

			//ignore [
			stack.Pop()
			var numString string
			// Find the number to make the final multiplier
			for stack.Length() > 0 && isCombinationOfNumbers(stack.Peek()) {
				numString = stack.Pop() + numString
			}
			num, _ := strconv.Atoi(numString)

			// create the final string k x string
			finalSubStr := strings.Repeat(str, num)
			stack.Push(finalSubStr)
			i++
		}
	}

	var finalString string
	for stack.Length() > 0 {
		finalString = stack.Pop() + finalString
	}
	return finalString
}

func isCombinationOfLetters(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func isCombinationOfNumbers(input string) bool {
	for _, char := range input {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func DecodeCaller() {

	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("3[abc]3[cd]ef"))
}
