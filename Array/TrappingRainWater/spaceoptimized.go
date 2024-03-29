package trappingrainwater

func Trap_optimized(height []int) int {
	volume := 0

	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0

	for left < right {

		if height[left] <= height[right] {

			if height[left] > leftMax {

				leftMax = height[left]
			} else {

				volume += (leftMax - height[left])
			}

			left++
		} else {

			if height[right] > rightMax {

				rightMax = height[right]
			} else {

				volume += (rightMax - height[right])
			}
			right--
		}
	}

	return volume
}

//0 1 0 2 0 3 3 0 4

// Another variant
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
