package graph

import("fmt")

// Lint code link->https://www.lintcode.com/problem/178/
func ValidTree(n int, edges [][]int) bool {

	// For a valid tree we :
	// 1) Should NOT have loops.
	// 2) Evenry node must be connected.

	adjMap := make(map[int][]int, 0)
	visitedMap := make(map[int]bool, 0)

	for i := 0; i < len(edges); i++ {

		source := edges[i][0]
		target := edges[i][1]

		adjMap[source] = append(adjMap[source], target)
		adjMap[target] = append(adjMap[target], source)
	}

	nodeCount := 0
	isValid := dfsCountNodes(0, -1, adjMap, visitedMap, &nodeCount)

	// Check if cycle is present
	if !isValid {

		return false
	}

	// Check if teh graph is connected
	if nodeCount != n {
		return false
	}

	return true
}

func dfsCountNodes(node int, prev int, adjList map[int][]int, visitedMap map[int]bool, count *int) bool {

	if visitedMap[node] {

		return false
	}

	visitedMap[node] = true
	*count++

	children := adjList[node]

	for _,child := range children {

		// Since it's an undirected graph, we don't want to return to the parent.
		if child == prev {
			continue
		}

		// Explicitly return false if we get a violation at any step without progressing forward.
		if dfsCountNodes(child, node, adjList, visitedMap, count) == false {

			return false
		}
	}

	return true
}

func ValidTreeCaller() {

	n:=5
	edges:=[][]int{
	{0,1},
	{0,2},
	{0,3},
	{1,4},
	}

	fmt.Print(ValidTree(n,edges))
}