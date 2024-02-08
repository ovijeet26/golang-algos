package revision

func trap(height []int) int {

	size := len(height)
	left, maxLeft, right, maxRight := 0, height[0], size-1, height[size-1]
	currentWater, totalWater := 0, 0

	for left < right {

		if maxLeft < maxRight {
			currentWater = maxLeft - height[left]
			left++

		} else {
			currentWater = maxRight - height[right]
			right--

		}

		if currentWater > 0 {
			totalWater += currentWater
		}
		maxLeft = maxHeight(height[left], maxLeft)
		maxRight = maxHeight(height[right], maxRight)
	}

	return totalWater
}

func maxHeight(a int, b int) int {

	if a < b {
		return b
	}

	return a
}
