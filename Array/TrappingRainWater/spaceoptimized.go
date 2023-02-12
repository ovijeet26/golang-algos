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
