package revision

func twoSum(nums []int, target int) []int {

	// key = number needed to make target | value = index of specific number
	hash := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {

		indexOfTarget, ok := hash[nums[i]]

		if ok {
			return []int{i, indexOfTarget}
		} else {

			tagetSum := target - nums[i]
			hash[tagetSum] = i
		}
	}

	return []int{}
}
