package graph

import "fmt"

// Leet code link -> https://leetcode.com/problems/course-schedule/
func canFinish(numCourses int, prerequisites [][]int) bool {

	// Utilize topological sort for traversing over graph and detect cycles.

	adjacencyMap := make(map[int][]int, 0)
	indegreeMap := make(map[int]int, numCourses)

	for i := 0; i < numCourses; i++ {

		adjacencyMap[i] = make([]int, 0)
	}

	for _, pair := range prerequisites {

		adjacencyMap[pair[1]] = append(adjacencyMap[pair[1]], pair[0])

		// Update indegree for the node where the edge is coming inwards.
		indegreeMap[pair[0]]++
	}

	// Execute topological sort algorithm to detect a cycle

	zeroStack := Stack{}
	visitedNodeCount := 0

	// Populate all zero indegree nodes in the stack.
	for i := 0; i < numCourses; i++ {

		if indegreeMap[i] == 0 {

			zeroStack.Push(i)
		}
	}

	for zeroStack.Size() != 0 {

		currentNode := zeroStack.Pop()
		visitedNodeCount++

		for _, child := range adjacencyMap[currentNode] {

			indegreeMap[child]--
			if indegreeMap[child] == 0 {
				zeroStack.Push(child)
			}
		}
	}

	if visitedNodeCount == numCourses {
		return true
	}

	return false
}

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {

	s.items = append(s.items, item)
}

func (s *Stack) Size() int {

	return len(s.items)
}

func (s *Stack) Pop() int {

	size := len(s.items)
	poppedItem := s.items[size-1]
	s.items = s.items[:size-1]
	return poppedItem
}

// // Alternate, less optimized solution.
// func canFinishBFS(numCourses int, prerequisites [][]int) bool {

// 	adjacencyMap := make(map[int][]int, 0)

// 	// Instantiate all the nodes in the adjacency map.
// 	for i := 0; i < numCourses; i++ {

// 		adjacencyMap[i] = make([]int, 0)
// 	}

// 	for _, pair := range prerequisites {

// 		adjacencyMap[pair[1]] = append(adjacencyMap[pair[1]], pair[0])
// 	}

// 	visited := struct{}{}

// 	toVisit := Queue{}

// 	for node := range adjacencyMap {

// 		// Perform BFS for every node in the list
// 		isVisited := make(map[int]struct{})
// 		toVisit.Enqueue(node)

// 		for toVisit.Size() != 0 {

// 			currentNode := toVisit.Dequeue().(int)
// 			isVisited[currentNode] = visited

// 			for _, child := range adjacencyMap[currentNode] {

// 				// if we detect a cycle, we can return false
// 				if child == node {
// 					return false
// 				}

// 				if _, ok := isVisited[child]; !ok {

// 					toVisit.Enqueue(child)
// 				}
// 			}
// 		}

// 	}

// 	return true
// }

// // Regular queue implementation

// type Queue struct {
// 	items []interface{}
// }

// func (q *Queue) Enqueue(data interface{}) {

// 	q.items = append(q.items, data)
// }

// func (q *Queue) Dequeue() interface{} {

// 	toRemove := q.items[0]

// 	// Remove first element from slice
// 	q.items = q.items[1:]
// 	return toRemove
// }

// func (q *Queue) Size() int {

// 	return len(q.items)
// }

func CanFinishCaller() {

	prerequisites := [][]int{
		{1, 0},
		{0, 1},
	}

	cF := canFinish(2, prerequisites)

	fmt.Print(cF)
}
