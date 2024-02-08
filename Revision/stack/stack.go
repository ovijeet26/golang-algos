package stack

type Stack struct {
	items []string
}

func (s *Stack) Len() int {

	return len(s.items)
}

func (s *Stack) Push(item string) {

	s.items = append(s.items, item)
}

func (s *Stack) Peek() string {

	return s.items[s.Len()-1]
}

func (s *Stack) Pop() string {

	size := s.Len()
	poppedItem := s.items[size-1]
	s.items = s.items[:size-1]
	return poppedItem
}

type IntStack struct {
	items []int
}

func (s *IntStack) Len() int {

	return len(s.items)
}

func (s *IntStack) Push(item int) {

	s.items = append(s.items, item)
}

func (s *IntStack) Peek() int {

	return s.items[s.Len()-1]
}

func (s *IntStack) Pop() int {

	size := s.Len()
	poppedItem := s.items[size-1]
	s.items = s.items[:size-1]
	return poppedItem
}

func (s *IntStack) ArrrayValues() []int {

	return s.items[]
}
