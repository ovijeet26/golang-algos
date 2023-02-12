package tree

import "fmt"

// Leet code link -> https://leetcode.com/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *TreeNode, k int) int {

	minHeap := MinHeap{data: new([]int)}
	dfsPopulateHeap(root, minHeap)

	kSmallest := 0
	for i := 0; i < k; i++ {

		kSmallest = minHeap.Extract()
	}

	return kSmallest
}

func dfsPopulateHeap(node *TreeNode, minHeap MinHeap) {

	if node == nil {
		return
	}

	minHeap.Insert(node.Value)

	dfsPopulateHeap(node.Left, minHeap)
	dfsPopulateHeap(node.Right, minHeap)
}

// Implementation of Min Heap data structure in Golang.
type MinHeap struct {
	data *[]int
}

// Add an element to the heap.
func (m *MinHeap) Insert(data int) {

	*m.data = append(*m.data, data)
	m.minHeapifyUp(len(*m.data) - 1)
}

// Peek the smallest key.
func (m *MinHeap) Peek() int {

	min := (*m.data)[0]
	return min
}

// Return the largest key, and remove it from the heap.
func (m *MinHeap) Extract() int {

	size := len(*m.data)

	max := (*m.data)[0]

	(*m.data)[0] = (*m.data)[size-1]
	*m.data = (*m.data)[:size-1]
	m.minHeapifyDown(0)

	return max
}

// Move the newly inserted data up the heap to it's proper position
func (m *MinHeap) minHeapifyUp(index int) {

	if index == 0 {
		return
	}

	for index > 0 && (*m.data)[index] < (*m.data)[m.parentIndex(index)] {

		m.swap(index, m.parentIndex(index))

		index = m.parentIndex(index)
	}
}

// Move the new root down the heap to it's proper position
func (m *MinHeap) minHeapifyDown(index int) {

	heapSize := len(*m.data)
	for (m.leftChildIndex(index) < heapSize && (*m.data)[m.leftChildIndex(index)] < (*m.data)[index]) ||
		(m.rightChildIndex(index) < heapSize && (*m.data)[m.rightChildIndex(index)] < (*m.data)[index]) {

		greaterIndex := index
		rightChildIndex := m.rightChildIndex(index)
		leftChildIndex := m.leftChildIndex(index)
		// If right child exists, then left child must exist as well.
		if rightChildIndex < heapSize && (*m.data)[rightChildIndex] < (*m.data)[leftChildIndex] {
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

	(*m.data)[i], (*m.data)[j] = (*m.data)[j], (*m.data)[i]
}

func KthElementCaller() {

	a := TreeNode{Value: 1}
	b := TreeNode{Value: 2}
	c := TreeNode{Value: 3}
	d := TreeNode{Value: 4}
	e := TreeNode{Value: 5}
	f := TreeNode{Value: 6}

	a.Left = &b
	a.Right = &c

	b.Left = &d
	c.Right = &e

	d.Left = &f

	kthSmallest := kthSmallest(&a, 3)

	fmt.Print(kthSmallest)

	//			1
	//		2		3
	//	4				5
	//6
}
