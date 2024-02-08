package revision

import "fmt"

func mergeSort(array []int, leftStart int, rightEnd int) {

	if leftStart >= rightEnd {
		return
	}

	mid := (leftStart + rightEnd) / 2
	mergeSort(array, leftStart, mid)
	mergeSort(array, mid+1, rightEnd)
	merge(array, leftStart, rightEnd)
}

func merge(array []int, leftStart int, rightEnd int) {

	left := leftStart
	right := rightEnd

	leftEnd := (leftStart + rightEnd) / 2
	rightStart := leftEnd + 1

	temp := make([]int, len(array))

	currentIndex := leftStart

	for leftStart <= leftEnd && rightStart <= rightEnd {

		if array[leftStart] <= array[rightStart] {
			temp[currentIndex] = array[leftStart]
			leftStart++
		} else {
			temp[currentIndex] = array[rightStart]
			rightStart++
		}
		currentIndex++
	}

	if leftStart > leftEnd {

		for rightStart <= rightEnd {
			temp[currentIndex] = array[rightStart]
			rightStart++
			currentIndex++
		}
	} else {

		for leftStart <= leftEnd {
			temp[currentIndex] = array[leftStart]
			currentIndex++
			leftStart++
		}
	}

	for i := left; i <= right; i++ {

		array[i] = temp[i]
	}
}

func MergeSortCaller() {

	nums := []int{4, 1, 2, -1, 6, 22, 33, 524, -234, 5, 0}

	fmt.Println(nums)
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
