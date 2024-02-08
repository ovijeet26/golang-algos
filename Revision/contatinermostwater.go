package revision

func maxArea(height []int) int {

	// in case we have only 0 or 1 elements, then no water can be contained
	if len(height) <= 1 {

		return 0
	}

	// if we have only 2 elements, then only the min height amount of water can be contained
	if len(height) == 2 {
		return min(height[0], height[1])
	}

	left := 0
	right := len(height) - 1

	maxArea := 0
	for left <= right {

		if left > len(height)-1 {
			break
		}

		if right < 0 {
			break
		}
		currentArea := area(left, height[left], right, height[right])

		if currentArea > maxArea {
			maxArea = currentArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func area(left int, leftHeight int, right int, rightHeight int) int {

	return min(leftHeight, rightHeight) * (right - left)
}

func min(a int, b int) int {

	if a < b {
		return a
	}
	return b
}
