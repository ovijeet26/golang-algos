package hard

import (
	"fmt"
	"strings"
)

// Alien dictionary
// https://practice.geeksforgeeks.org/problems/alien-dictionary/1
func findOrderString(words []string, n int, k int) string {

	adjList := make(map[string]map[string]bool, k)
	indegreeMap := make(map[string]int, 0)

	// Create adjacency list and initialize an indegree map
	for i := 0; i < n-1; i++ {

		firstWord := words[i]
		secondWord := words[i+1]

		i := 0
		firstWordSize := len(firstWord)
		sencondWordSize := len(secondWord)

		// In case the two words have with the same character at same position,
		// we iterate to find the first difference
		for firstWord[i] == secondWord[i] && firstWordSize > i && sencondWordSize > i {
			i++
		}

		// We append the list, only if we have a valid difference and both strings are still eft to be traversed.
		if firstWordSize > i && sencondWordSize > i {

			if adjList[string(firstWord[i])] == nil {

				adjList[string(firstWord[i])] = map[string]bool{}
				indegreeMap[string(firstWord[i])] = 0
			}

			adjList[string(firstWord[i])][string(secondWord[i])] = true
		}
	}

	// Populate indegrree map.
	for _, edges := range adjList {

		for child := range edges {

			indegreeMap[child]++
		}
	}

	// Run topo sort to print order of node
	zeroStack := Stack{}
	visitedNodeCount := 0
	answerArray := make([]string, k)

	// Populate all zero indegree nodes in the stack.
	for node, indegree := range indegreeMap {

		if indegree == 0 {
			zeroStack.Push(node)
		}
	}

	ansIndex := 0
	for zeroStack.Size() != 0 {

		currentNode := zeroStack.Pop()
		visitedNodeCount++

		answerArray[ansIndex] = currentNode
		ansIndex++
		for child := range adjList[currentNode] {

			indegreeMap[child]--
			if indegreeMap[child] == 0 {
				zeroStack.Push(child)
			}
		}
	}

	// If there is no cycle
	if visitedNodeCount == k {
		return strings.Join(answerArray, "")
	}

	return ""
}

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {

	s.items = append(s.items, item)
}

func (s *Stack) Size() int {

	return len(s.items)
}

func (s *Stack) Pop() string {

	size := len(s.items)
	poppedItem := s.items[size-1]
	s.items = s.items[:size-1]
	return poppedItem
}

func AlienCaller() {

	dict := []string{"baa", "abcd", "abca", "cab", "cad"}
	n := 5
	k := 4

	// n := 3
	// k := 3
	// dict := []string{"caa", "aaa", "aab"}

	fmt.Print(findOrderString(dict, n, k))
}
