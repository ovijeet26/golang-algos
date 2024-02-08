package array

func twoSum(numbers []int, target int) []int {

	left := 0
	right := len(numbers) - 1

	for left <= right {

		if numbers[left]+numbers[right] == target {

			//+1 because the problem states to return 1 indexed values
			return []int{left + 1, right + 1}
		}

		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
