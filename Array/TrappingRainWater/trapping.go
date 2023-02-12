package trappingrainwater

func Trap(height []int) int {
	volume := 0

	leftMaxHeightAtIndex := make([]int, 0)
	rightMaxHeightAtIndex := make([]int, 0)
	maxHeight := 0

	for pointer := 0; pointer < len(height); pointer++ {

		if height[pointer] > maxHeight {

			maxHeight = height[pointer]
		}
		leftMaxHeightAtIndex[pointer] = maxHeight
	}

	maxHeight = 0

	for pointer := len(height) - 1; pointer >= 0; pointer-- {

		if height[pointer] > maxHeight {

			maxHeight = height[pointer]
		}
		rightMaxHeightAtIndex[pointer] = maxHeight
	}

	for pointer := 0; pointer < len(height); pointer++ {

		volumeAtIndex := min(leftMaxHeightAtIndex[pointer], rightMaxHeightAtIndex[pointer]) - height[pointer]

		if volumeAtIndex > 0 {

			volume += volumeAtIndex
		}
	}

	return volume
}

func min(a int, b int) int {

	if a < b {
		return a
	}
	return b
}

//0 1 0 2 0 1
