package sort

func QuickSort(data []int, left int, right int) {

	if left < right {

		partitionIndex := partition(data, left, right)
		QuickSort(data, left, partitionIndex-1)
		QuickSort(data, partitionIndex+1, right)
	}
}

func partition(data []int, left int, right int) int {

	pivot := data[right]
	partitionIndex := left

	for j := left; j < right; j++ {

		if data[j] < pivot {

			swap(data, j, partitionIndex)
			partitionIndex++
		}
	}

	swap(data, partitionIndex, right)
	return partitionIndex
}

func swap(data []int, i int, j int) {

	temp := data[i]
	data[i] = data[j]
	data[j] = temp
}
