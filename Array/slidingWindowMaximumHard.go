package array

import "fmt"

// Leet code link -> https://leetcode.com/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {

	finalArray := make([]int, 0)
	indexDeq := Deque{}

	leftIndex := 0
	rightIndex := 0

	for rightIndex = 0; rightIndex < len(nums); rightIndex++ {

		// Remove items from Deque which are lesser than the current value
		for indexDeq.Size() > 0 && nums[indexDeq.PeekRight()] < nums[rightIndex] {

			indexDeq.PopRight()
		}

		// Push the current element into the Deque
		indexDeq.Push(rightIndex)

		// Add the max value (Leftmost element of Deque) to answer array when the sliding window length is met.
		if rightIndex+1 >= k {

			finalArray = append(finalArray, nums[indexDeq.PeekLeft()])
			leftIndex++
		}

		// Remove leftmost element from Deque when the leftIndex is greater than that index.
		if leftIndex > indexDeq.PeekLeft() {

			indexDeq.PopLeft()
		}
	}

	return finalArray
}

type Deque struct {
	items []int
}

func (d *Deque) Push(data int) {

	d.items = append(d.items, data)
}

func (d *Deque) PopLeft() int {

	item := d.items[0]

	d.items = d.items[1:]

	return item
}

func (d *Deque) PeekLeft() int {

	return d.items[0]
}

func (d *Deque) PopRight() int {

	item := d.items[d.Size()-1]

	d.items = d.items[:d.Size()-1]

	return item
}

func (d *Deque) PeekRight() int {

	return d.items[d.Size()-1]
}

func (d *Deque) Size() int {

	return len(d.items)
}

func SlidingWindowMaxCaller() {

	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	// nums := []int{7, 2, 4}
	// k := 2

	fmt.Print(maxSlidingWindow(nums, k))
}
