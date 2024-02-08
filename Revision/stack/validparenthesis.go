package stack

func isValid(s string) bool {

	bracketMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	bStack := Stack{}
	for _, char := range s {

		current := string(char)
		bracket, ok := bracketMap[current]

		if ok {
			bStack.Push(bracket)
		} else {

			if bStack.Len() <= 0 {
				return false
			}

			if current != bStack.Pop() {
				return false
			}
		}
	}

	if bStack.Len() > 0 {
		return false
	}

	return true
}

func ValidCaller() {

	isValid("()")
}
