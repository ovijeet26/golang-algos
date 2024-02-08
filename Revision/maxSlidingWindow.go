package revision

func maxSlidingWindow(nums []int, k int) []int {

	maxDeq := Deque{}
	left, right := 0, 0
	output := []int{}

	for right < len(nums) {

		for maxDeq.Size() > 0 && nums[maxDeq.PeekRight()] < nums[right] {
			maxDeq.PopRight()
		}

		maxDeq.Push(right)

		if right+1 >= k {
			output = append(output, nums[maxDeq.PeekLeft()])
			left++
		}

		right++

		if left > maxDeq.PeekLeft() {
			maxDeq.PopLeft()
		}
	}
	return output
}
