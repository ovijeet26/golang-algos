package graph

import (
	"fmt"
	"math"
)

// Leet code link -> https://leetcode.com/problems/network-delay-time/
// You are given a network of n nodes, labeled from 1 to n. You are also given times,
// a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node,
// vi is the target node, and wi is the time it takes for a signal to travel from source to target.
func networkDelayTime(times [][]int, n int, k int) int {

	// Solve using Dijkstra's algorithm. Note: Dijkstra's algorithm does not account fopr negative weights.
	// We assume that there is no negative wights in this graph.
	adjacencyMap := make(map[int][]weightedNode)
	distances := make([]int, n)

	infinity := int(math.Inf(1))

	// Initialize all distances with positive infinity value.
	for i := 0; i < n; i++ {

		distances[i] = infinity
	}

	for _, time := range times {

		// All are 0 indexed, hence we will subtract 1
		source := time[0] - 1
		target := time[1] - 1
		weight := time[2]

		adjacencyMap[source] = append(adjacencyMap[source], weightedNode{node: target, weight: weight})
	}

	// Use a priority queue or min heap, to track the minumum from the array to perform dijkstra's algorithm.
	priorityQueue := MinHeap{}

	// Make the distance of the source node (k) as 0. Since arrays are 0 indexed, we use k-1.
	distances[k-1] = 0
	priorityQueue.Insert(k-1, distances)

	for priorityQueue.Size() > 0 {

		currentNode := priorityQueue.Extract(distances)
		weightTillCurrentNode := distances[currentNode]
		adjascentNodes := adjacencyMap[currentNode]

		for _, childNode := range adjascentNodes {

			weightTillNode := weightTillCurrentNode + childNode.weight

			if weightTillNode < distances[childNode.node] {

				distances[childNode.node] = weightTillNode
				priorityQueue.Insert(childNode.node, distances)

			}

		}
	}

	max := 0

	// Calculate the max distance in the array, that will be the longest time to reach all nodes.
	for _, distance := range distances {

		if distance > max {
			max = distance
		}
	}

	if max == infinity {

		return -1
	}
	return max
}

type weightedNode struct {
	node   int
	weight int
}

// Implementation of Min Heap data structure in Golang.
type MinHeap struct {
	data []int
}

// Size of the heap.
func (m *MinHeap) Size() int {

	return len(m.data)
}

// Add an element to the heap.
func (m *MinHeap) Insert(data int, distances []int) {

	m.data = append(m.data, data)
	m.maxHeapifyUp(len(m.data)-1, distances)
}

// Return the largest key, and remove it from the heap.
func (m *MinHeap) Extract(distances []int) int {

	size := len(m.data)

	min := m.data[0]

	m.data[0] = m.data[size-1]
	m.data = m.data[:size-1]
	m.maxHeapifyDown(0, distances)

	return min
}

// Move the newly inserted data up the heap to it's proper position
func (m *MinHeap) maxHeapifyUp(index int, distances []int) {

	if index == 0 {
		return
	}

	for index > 0 && distances[m.data[index]] < distances[m.data[m.parentIndex(index)]] {

		m.swap(index, m.parentIndex(index))

		index = m.parentIndex(index)
	}
}

// Move the new root down the heap to it's proper position
func (m *MinHeap) maxHeapifyDown(index int, distances []int) {

	heapSize := len(m.data)
	for (m.leftChildIndex(index) < heapSize && distances[m.data[m.leftChildIndex(index)]] < distances[m.data[index]]) ||
		(m.rightChildIndex(index) < heapSize && distances[m.data[m.rightChildIndex(index)]] < distances[m.data[index]]) {

		lesserIndex := index
		rightChildIndex := m.rightChildIndex(index)
		leftChildIndex := m.leftChildIndex(index)
		// If right child exists, then left child must exist as well.
		if rightChildIndex < heapSize && distances[m.data[rightChildIndex]] < distances[m.data[leftChildIndex]] {
			lesserIndex = rightChildIndex
		} else {
			lesserIndex = leftChildIndex
		}

		m.swap(index, lesserIndex)
		index = lesserIndex
	}

}

func (m *MinHeap) parentIndex(index int) int {

	return (index - 1) / 2
}

func (m *MinHeap) leftChildIndex(index int) int {

	return (2 * index) + 1
}

func (m *MinHeap) rightChildIndex(index int) int {

	return (2 * index) + 2
}

func (m *MinHeap) swap(i int, j int) {

	//m.data[i], m.data[j] = m.data[j], m.data[i]
	temp := m.data[i]
	m.data[i] = m.data[j]
	m.data[j] = temp
}

func NetworkDelayCaller() {

	times := [][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
	}
	n := 4
	k := 2

	// times := [][]int{{1,2,1}}
	// n := 2
	// k := 2

	// times := [][]int{{1,2,1}}
	// n := 2
	// k := 1

	nTime := networkDelayTimeBf(times, n, k)

	fmt.Print(nTime)
}
