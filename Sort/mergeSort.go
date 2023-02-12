package sort

func MergeSort(data []int, start int, end int) {

	if start < end {
		mid := (start + end) / 2
		MergeSort(data, start, mid)
		MergeSort(data, mid+1, end)
		merge(data, start, mid, end)
	}
}

func merge(data []int, start int, mid int, end int) {

	tempArr := make([]int, end-start+1)

	i, j := start, mid+1
	counter := 0

	for i <= mid && j <= end {

		if data[i] <= data[j] {

			tempArr[counter] = data[i]
			i++
			counter++
		} else {

			tempArr[counter] = data[j]
			j++
			counter++
		}
	}

	for i <= mid {

		tempArr[counter] = data[i]
		i++
		counter++
	}

	for j <= end {

		tempArr[counter] = data[j]
		j++
		counter++
	}

	counter = 0
	for p := start; p <= end; p++ {

		data[p] = tempArr[counter]
		counter++
	}
}
