package heap

// Implementation of Max Heap data structure in Golang.
type MaxHeap struct {
	data []int
}

// Add an element to the heap.
func (m *MaxHeap) Insert(data int) {

	m.data = append(m.data, data)
	m.maxHeapifyUp(len(m.data) - 1)
}

// Peek the largest key.
func (m *MaxHeap) Peek() int {

	max := m.data[0]
	return max
}

// Return the largest key, and remove it from the heap.
func (m *MaxHeap) Extract() int {

	size := len(m.data)

	max := m.data[0]

	m.data[0] = m.data[size-1]
	m.data = m.data[:size-1]
	m.maxHeapifyDown(0)

	return max
}

// Move the newly inserted data up the heap to it's proper position
func (m *MaxHeap) maxHeapifyUp(index int) {

	if index == 0 {
		return
	}

	for index > 0 && m.data[index] > m.data[m.parentIndex(index)] {

		m.swap(index, m.parentIndex(index))

		index = m.parentIndex(index)
	}
}

// Move the new root down the heap to it's proper position
func (m *MaxHeap) maxHeapifyDown(index int) {

	heapSize := len(m.data)
	for (m.leftChildIndex(index) < heapSize && m.data[m.leftChildIndex(index)] > m.data[index]) ||
		(m.rightChildIndex(index) < heapSize && m.data[m.rightChildIndex(index)] > m.data[index]) {

		greaterIndex := index
		rightChildIndex := m.rightChildIndex(index)
		leftChildIndex := m.leftChildIndex(index)
		// If right child exists, then left child must exist as well.
		if rightChildIndex < heapSize && m.data[rightChildIndex] > m.data[leftChildIndex] {
			greaterIndex = rightChildIndex
		} else {
			greaterIndex = leftChildIndex
		}

		m.swap(index, greaterIndex)
		index = greaterIndex
	}

}

func (m *MaxHeap) parentIndex(index int) int {

	return (index - 1) / 2
}

func (m *MaxHeap) leftChildIndex(index int) int {

	return (2 * index) + 1
}

func (m *MaxHeap) rightChildIndex(index int) int {

	return (2 * index) + 2
}

func (m *MaxHeap) swap(i int, j int) {

	//m.data[i], m.data[j] = m.data[j], m.data[i]
	temp := m.data[i]
	m.data[i] = m.data[j]
	m.data[j] = temp
}
