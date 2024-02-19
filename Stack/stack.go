package stack

import "strings"

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {

	s.items = append(s.items, item)
}

func (s *Stack) Length() int {

	return len(s.items)
}

func (s *Stack) Peek() string {

	return s.items[s.Length()-1]
}

func (s *Stack) Pop() string {

	size := len(s.items)
	poppedItem := s.items[size-1]
	s.items = s.items[:size-1]
	return poppedItem
}

func (s *Stack) ArrayValues() []string {
	return s.items
}

// Leet code link - https://leetcode.com/problems/valid-parentheses/submissions/
func isValid(s string) bool {

	//bracketMap := make(map[string]string,0)

	bracketMap := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	bracketStack := Stack{}
	length := len(s)

	for i := 0; i <= length-1; i++ {

		currentBracket := string(s[i])

		bracket, ok := bracketMap[currentBracket]

		if ok {

			bracketStack.Push(bracket)
		} else {

			if bracketStack.Length() == 0 {

				return false
			}

			poppedBracket := bracketStack.Pop()

			if poppedBracket != string(s[i]) {

				return false
			}
		}
	}

	if bracketStack.Length() != 0 {

		return false
	}
	return true
}

// e.g. a)b(c)d, ))()((
// Leet code link -> https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
func minRemoveToMakeValid(str string) string {

	bracketStack := IntStack{}
	length := len(str)

	s := strings.Split(str, "")
	// Traverse the entire string
	for i := 0; i < length; i++ {

		// If we get an open bracket, then push the index in the stack
		if string(s[i]) == "(" {

			// Push the index of the opening bracket
			bracketStack.Push(i)
		}
		// If we get a close bracket, then pop an equivalent open bracket from the stack
		if string(s[i]) == ")" {

			// If we don’t have any opening bracket in the stack, then we can directly delete the closing bracket. Else pop an equivalent opening bracket.
			if bracketStack.Size() == 0 {
				s[i] = ""
			} else {

				bracketStack.Pop()
			}
		}

	}

	remainingStackSize := bracketStack.Size()

	for i := 0; i < remainingStackSize; i++ {

		pos := bracketStack.Pop()
		s[pos] = ""
	}

	return strings.Join(s, "")
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

func (s *IntStack) ArrayValues() []int {
	return s.items
}

type GenericStack struct {
	items []interface{}
}

func (s *GenericStack) Push(value interface{}) {

	s.items = append(s.items, value)
}
func (s *GenericStack) Pop() interface{} {

	size := len(s.items)
	poppedItem := s.items[size-1]
	s.items = s.items[:size-1]
	return poppedItem
}

func (s *GenericStack) Size() int {

	return len(s.items)
}

func (s *GenericStack) Peek() interface{} {

	return s.items[s.Size()-1]
}
