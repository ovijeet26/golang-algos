package twosum

func TwoSum(nums []int, target int) []int {

	var resultIndices []int
	valueIndexMap := make(map[int]int)
	var numberToFind int

	// Populate the value index map with number to find and index values
	for index, value := range nums {

		indexOfTarget, ok := valueIndexMap[value]

		if ok {

			resultIndices = append(resultIndices, indexOfTarget, index)
			return resultIndices
		} else {

			numberToFind = target - value
			valueIndexMap[numberToFind] = index
		}
	}

	return nil
}
