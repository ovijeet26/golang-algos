package revision

func longestSubarray(nums []int, limit int) int {

	maxDeq, minDeq := Deque{}, Deque{}
	maxCount := 0

	left, right := 0, 0
	for right = 0; right < len(nums); right++ {

		for maxDeq.Size() > 0 && nums[right] > nums[maxDeq.PeekRight()] {
			maxDeq.PopRight()
		}

		for minDeq.Size() > 0 && nums[right] < nums[minDeq.PeekRight()] {
			minDeq.PopRight()
		}

		maxDeq.Push(right)
		minDeq.Push(right)

		absoluteDifference := Absolute(nums[maxDeq.PeekLeft()] - nums[minDeq.PeekLeft()])
		if absoluteDifference > limit {

			if maxDeq.PeekLeft() <= left {
				maxDeq.PopLeft()
			}

			if minDeq.PeekLeft() <= left {
				minDeq.PopLeft()
			}
			left++
			continue
		}

		currentWindowSize := (right - left) + 1
		if currentWindowSize > maxCount {
			maxCount = currentWindowSize
		}
	}

	return maxCount
}

func Absolute(a int) int {

	if a < 0 {
		return -1
	}

	return a
}

func LongestCaller() {

	nums := []int{4, 2, 2, 2, 4, 4, 2, 2}

	print(longestSubarray(nums, 0))
}
