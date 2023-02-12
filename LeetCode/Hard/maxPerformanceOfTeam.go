package hard

import (
	"fmt"
	"math"
)

// Leet code link -> https://leetcode.com/problems/maximum-performance-of-a-team/
func maxPerformance(n int, speed []int, efficiency []int, k int) int {

	quickSort(efficiency, speed)

	minSpeedHeap := MinHeap{}
	maxPerformance := 0
	cumulativeSpeed := 0

	for i := 0; i < n; i++ {

		minSpeedHeap.Insert(speed[i])
		if i < k {

			cumulativeSpeed += speed[i]
		} else {

			cumulativeSpeed += speed[i]

			cumulativeSpeed -= minSpeedHeap.Extract()
		}

		performance := efficiency[i] * cumulativeSpeed

		if performance > maxPerformance {

			maxPerformance = performance
		}
	}

	return maxPerformance % (int(math.Pow(10, 9) + 7))
}

func quickSort(eff []int, speed []int) {

	qs(eff, 0, len(eff)-1, speed)
}

func qs(eff []int, left int, right int, speed []int) {

	if left < right {

		partitionKey := partition(eff, left, right, speed)
		qs(eff, left, partitionKey-1, speed)
		qs(eff, partitionKey+1, right, speed)
	}
}

func partition(eff []int, left int, right int, speed []int) int {

	pivot := eff[right]
	partitionIndex := left

	for i := left; i < right; i++ {

		if eff[i] > pivot {
			swap(eff, partitionIndex, i, speed)
			partitionIndex++
		}
	}

	swap(eff, partitionIndex, right, speed)
	return partitionIndex
}

func swap(eff []int, a int, b int, speed []int) {

	eff[a], eff[b] = eff[b], eff[a]
	speed[a], speed[b] = speed[b], speed[a]
}

// Implementation of Max Heap data structure in Golang.
type MinHeap struct {
	data []int
}

// Add an element to the heap.
func (m *MinHeap) Insert(data int) {

	m.data = append(m.data, data)
	m.minHeapifyUp(len(m.data) - 1)
}

// Return the largest key, and remove it from the heap.
func (m *MinHeap) Extract() int {

	size := len(m.data)

	max := m.data[0]

	m.data[0] = m.data[size-1]
	m.data = m.data[:size-1]
	m.minHeapifyDown(0)

	return max
}

// Move the newly inserted data up the heap to it's proper position
func (m *MinHeap) minHeapifyUp(index int) {

	if index == 0 {
		return
	}

	for index > 0 && m.data[index] < m.data[m.parentIndex(index)] {

		m.swap(index, m.parentIndex(index))

		index = m.parentIndex(index)
	}
}

// Move the new root down the heap to it's proper position
func (m *MinHeap) minHeapifyDown(index int) {

	heapSize := len(m.data)
	for (m.leftChildIndex(index) < heapSize && m.data[m.leftChildIndex(index)] < m.data[index]) ||
		(m.rightChildIndex(index) < heapSize && m.data[m.rightChildIndex(index)] < m.data[index]) {

		greaterIndex := index
		rightChildIndex := m.rightChildIndex(index)
		leftChildIndex := m.leftChildIndex(index)
		// If right child exists, then left child must exist as well.
		if rightChildIndex < heapSize && m.data[rightChildIndex] < m.data[leftChildIndex] {
			greaterIndex = rightChildIndex
		} else {
			greaterIndex = leftChildIndex
		}

		m.swap(index, greaterIndex)
		index = greaterIndex
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

func MaxPerformanceCaller() {

	n := 6
	speed := []int{2, 10, 3, 1, 5, 8}
	efficiency := []int{5, 4, 3, 9, 7, 2}
	k := 2

	fmt.Print(maxPerformance(n, speed, efficiency, k))
}

// s[1,2,3,4]
// e[9,8,7,6]

// 1*9

// 1*8 + 2*8  -> (1+2)*8

// 1*7 + 2*7 + 3*7 -> (1+2+3)*7
