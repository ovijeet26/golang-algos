package revision

func maxSubArray(nums []int) int {

	current, max := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] > (nums[i] + current) {
			current = nums[i]
		} else {
			current += nums[i]
		}

		if max < current {
			max = current
		}
	}

	return max
}
