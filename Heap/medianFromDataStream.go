package heap

// Leet code link -> 
type MedianFinder struct {
	leftHeap  MaxHeap
	rightHeap MinHeap
	leftSize  int
	rightSize int
}

func Constructor() MedianFinder {

	return MedianFinder{
		leftHeap:  MaxHeap{},
		rightHeap: MinHeap{},
		leftSize:  0,
		rightSize: 0,
	}
}

func (this *MedianFinder) AddNum(num int) {

	this.leftHeap.Insert(num)
	this.leftSize++

	sizeDifference := this.leftSize - this.rightSize

	if sizeDifference > 1 {

		// Push to other heap
		extraNum := this.leftHeap.Extract()
		this.leftSize--
		this.rightHeap.Insert(extraNum)
		this.rightSize++
	}

	if this.leftSize > 0 && this.rightSize > 0 && this.leftHeap.Peek() > this.rightHeap.Peek() {

		// Swap
		left := this.leftHeap.Extract()
		right := this.rightHeap.Extract()

		this.leftHeap.Insert(right)
		this.rightHeap.Insert(left)
	}

}

func (this *MedianFinder) FindMedian() float64 {

	if this.leftSize == this.rightSize {

		left := float64(this.leftHeap.Peek())
		right := float64(this.rightHeap.Peek())

		return (left + right) / 2
	}

	if this.leftSize > this.rightSize {

		return float64(this.leftHeap.Peek())
	}

	return float64(this.rightHeap.Peek())
}
