package containerwithmostwater

func maxArea(height []int) int {

	maxArea, p1, p2 := 0, 0, len(height)-1

	for p1 < p2 {

		area := min(height[p1], height[p2]) * (p2 - p1)
		if area > maxArea {
			maxArea = area
		}

		// Increment or decrement the pointer based on the value of the height
		if height[p1] > height[p2] {
			p2--
		} else {
			p1++
		}

	}

	return maxArea
}

// Find the minimum between two numbers
func min(a int, b int) int {

	if a < b {
		return a
	}
	return b
}
