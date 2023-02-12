package sort

func QuickSelect(data []int, left int, right int, indexToFind int) {

	if left < right {

		partitionIndex := partitionData(data, left, right)

		if partitionIndex == indexToFind {
			return
		} else if partitionIndex > indexToFind {
			QuickSelect(data, left, partitionIndex-1, indexToFind)
		} else {
			QuickSelect(data, partitionIndex+1, right, indexToFind)
		}
	}
}

func partitionData(data []int, left int, right int) int {

	pivot := data[right]
	partitionIndex := left

	for j := left; j < right; j++ {

		if data[j] < pivot {
			swapData(data, partitionIndex, j)
			partitionIndex++
		}
	}
	swapData(data, partitionIndex, right)
	return partitionIndex
}

func swapData(data []int, i int, j int) {

	//data[i], data[j] = data[j], data[i]
	temp := data[i]
	data[i] = data[j]
	data[j] = temp
}
