package graph

import "fmt"

// Leet code link -> https://leetcode.com/problems/time-needed-to-inform-all-employees/
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {

	adjacencyMap := make(map[int][]int, 0)

	for index, manager := range manager {

		// We don't want to include the level above the manger into the node list in the adjacency list
		if manager == -1 {
			continue
		}

		adjacencyMap[manager] = append(adjacencyMap[manager], index)
	}

	return mostTimeDfs(headID, adjacencyMap, informTime)
}

func mostTimeDfs(currentNode int, adjacencyMap map[int][]int, informTime []int) int {

	if _, ok := adjacencyMap[currentNode]; !ok {

		return 0
	}

	maxTime := 0

	for _, child := range adjacencyMap[currentNode] {

		maxTime = max(maxTime, mostTimeDfs(child, adjacencyMap, informTime))
	}

	return maxTime + informTime[currentNode]
}

func max(a int, b int) int {

	if a >= b {
		return a
	}
	return b
}

func NumOfMinutesCaller() {

	n := 6
	headID := 2
	manager := []int{2, 2, -1, 2, 2, 2}
	informTime := []int{0, 0, 1, 0, 0, 0}

	fmt.Println(numOfMinutes(n, headID, manager, informTime))
}
